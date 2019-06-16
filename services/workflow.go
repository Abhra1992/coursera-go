package services

import (
	"coursera/types"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type IWorkflow interface {
	DownloadModules(modules []*types.Module) (bool, error)
}

type CourseraWorkflow struct {
	scheduler IDownloadScheduler
	args      *types.Arguments
	className string
}

func NewCourseraWorkflow(dw IDownloadScheduler, args *types.Arguments, className string) *CourseraWorkflow {
	return &CourseraWorkflow{dw, args, className}
}

func (cw *CourseraWorkflow) DownloadModules(modules []*types.Module) (bool, error) {
	xpath := "."
	if cw.args.Path != "" {
		xpath = cw.args.Path
	}
	cpath := filepath.Join(xpath, cw.className)
	if err := ensureDirExists(cpath); err != nil {
		return false, err
	}
	for _, module := range modules {
		log.Printf("MODULE %s", module.Name)
		for _, section := range module.Sections {
			log.Printf("\tSECTION %s", section.Name)
			spath := filepath.Join(cpath, module.Symbol, section.Symbol)
			if err := ensureDirExists(spath); err != nil {
				return false, err
			}
			for ii, item := range section.Items {
				log.Printf("\t\t%s ITEM %s", item.Type, item.Symbol)
				for ext, link := range item.Links {
					fname := filepath.Join(spath, fmt.Sprintf("%02d-%s.%s", ii, item.Symbol, ext))
					cw.scheduler.Download(link, fname)
				}
			}
		}
	}
	return true, nil
}

func ensureDirExists(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

func (cw *CourseraWorkflow) HandleResource() {

}
