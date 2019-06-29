package services

import (
	"coursera/api"
	"coursera/scheduler"
	"coursera/types"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CourseraWorkflow sets up the workflow for downloading Coursera class resources
type CourseraWorkflow struct {
	scheduler   scheduler.IScheduler
	args        *types.Arguments
	className   string
	skippedURLs []string
	failedURLs  []string
}

// NewCourseraWorkflow constructor
func NewCourseraWorkflow(dw scheduler.IScheduler, args *types.Arguments, className string) *CourseraWorkflow {
	return &CourseraWorkflow{dw, args, className, make([]string, 0), make([]string, 0)}
}

// DownloadModules downloads the modules in the Coursera class
func (cw *CourseraWorkflow) DownloadModules(modules []*types.Module) (bool, error) {
	_, cpath, err := cw.resolveEnsureExecutionPaths()
	if err != nil {
		return false, err
	}
	for _, module := range modules {
		log.Printf("MODULE %s", module.Name)
		lastUpdate := time.Time{}
		for _, section := range module.Sections {
			log.Printf("\tSECTION %s", section.Name)
			spath := filepath.Join(cpath, module.Symbol, section.Symbol)
			if err := EnsureDirExists(spath); err != nil {
				return false, err
			}
			for ii, item := range section.Items {
				log.Printf("\t\t%s ITEM %s", item.Type, item.Symbol)
				for _, res := range item.Resources {
					fname := filepath.Join(spath, fmt.Sprintf("%02d-%s.%s", ii, item.Symbol, res.Extension))
					cw.handleResource(res.Link, res.Extension, fname, lastUpdate)
				}
			}
		}
	}
	cw.scheduler.Complete()
	return true, err
}
func (cw *CourseraWorkflow) resolveEnsureExecutionPaths() (string, string, error) {
	xpath := "."
	if cw.args.Path != "" {
		xpath = cw.args.Path
	}
	cpath := filepath.Join(xpath, cw.className)
	if err := EnsureDirExists(cpath); err != nil {
		return "", "", err
	}
	return xpath, cpath, nil
}

func (cw *CourseraWorkflow) handleResource(link string, format string, fname string, lastUpdate time.Time) (time.Time, error) {
	overwrite, resume, skipDownload := cw.args.Overwrite, cw.args.Resume, cw.args.SkipDownload
	exists, err := FileExists(fname)
	if err != nil {
		return lastUpdate, err
	}
	if overwrite || resume || !exists {
		if !skipDownload {
			if strings.HasPrefix(link, api.InMemoryMarker) {
				pageContent := strings.TrimPrefix(link, api.InMemoryMarker)
				log.Printf("Saving page contents to: %s", fname)
				ioutil.WriteFile(fname, []byte(pageContent), 0644)
			} else if cw.skippedURLs != nil && shouldSkipFormatURL(format, link) {
				cw.skippedURLs = append(cw.skippedURLs, link)
			} else {
				dt := scheduler.DownloadTask{URL: link, File: fname, Callback: cw.onCompletionCallback}
				cw.scheduler.Schedule(dt)
			}
		} else {
			// touch file
			f, err := os.OpenFile(fname, os.O_CREATE, 0644)
			if err != nil {
				log.Printf("Could not touch file [%s]", fname)
			}
			f.Close()
		}
		lastUpdate = time.Now()
	} else {
		log.Printf("\t\t> Exists [%s]", fname)
		fi, err := os.Stat(fname)
		if err != nil {
			return lastUpdate, err
		}
		mtime := fi.ModTime()
		if mtime.After(lastUpdate) {
			lastUpdate = mtime
		}
	}
	return lastUpdate, nil
}

func (cw *CourseraWorkflow) onCompletionCallback(link string, err error) {
	if err != nil {
		log.Println(err.Error())
		cw.failedURLs = append(cw.failedURLs, link)
	}
}

func (cw *CourseraWorkflow) runHooks() {}
