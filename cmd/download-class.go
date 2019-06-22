package cmd

import (
	"coursera/api"
	"coursera/services"
	"coursera/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func DownloadOnDemandClass(cs *api.CourseraSession, className string, args *types.Arguments) (bool, error) {
	extractor := services.NewCourseraExtractor(cs, args)
	// Check if syllabus is cached - if yes, use it
	sf := fmt.Sprintf("%s-syllabus.json", className)
	exists, err := fileExists(sf)
	if err != nil {
		log.Printf("Error when checking for existence of syllabus for %s", className)
		return false, err
	}
	var modules []*types.Module
	if args.CacheSyllabus && exists {
		syl, err := ioutil.ReadFile(sf)
		if err != nil {
			log.Printf("Could not read syllabus for %s", className)
			return false, err
		}
		json.Unmarshal(syl, &modules)
	} else {
		modules, err = extractor.GetModules(className)
		if err != nil {
			fmt.Println("Error getting Modules")
		}
	}
	if args.CacheSyllabus {
		jsyl, err := json.Marshal(modules)
		if err != nil {
			log.Printf("Could not cache syllabus for %s", className)
			return false, err
		}
		ioutil.WriteFile(sf, jsyl, 0644)
	}
	if args.OnlySyllabus {
		return true, nil
	}
	downloader := services.GetDownloader(cs)
	scheduler := services.NewConsecutiveScheduler(downloader)
	workflow := services.NewCourseraWorkflow(scheduler, args, className)
	completed, err := workflow.DownloadModules(modules)
	return completed, err
}

func fileExists(fname string) (bool, error) {
	_, err := os.Stat(fname)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
