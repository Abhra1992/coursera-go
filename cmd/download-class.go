package cmd

import (
	"coursera/api"
	"coursera/services"
	"coursera/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// DownloadOnDemandClass downloads a single Coursera On Demand class
func DownloadOnDemandClass(cs *api.CourseraSession, className string, args *types.Arguments) (bool, error) {
	extractor := services.NewCourseraExtractor(cs, args)
	var modules []*types.Module
	// Check if syllabus is cached - if yes, use it
	sf := fmt.Sprintf("%s-syllabus.json", className)
	if args.CacheSyllabus {
		syllabusExists, err := services.FileExists(sf)
		if err != nil {
			log.Printf("Error when checking for existence of syllabus for %s", className)
			return false, err
		}
		if syllabusExists {
			syl, err := ioutil.ReadFile(sf)
			if err != nil {
				log.Printf("Could not read syllabus for %s", className)
				return false, err
			}
			json.Unmarshal(syl, &modules)
		}
	}
	if modules == nil {
		ems, err := extractor.GetModules(className)
		modules = ems
		if err != nil {
			fmt.Println("Error getting Modules")
		}
	}
	// Check if syllabus should be cached - if yes, save it
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
	downloader := services.GetDownloader(cs, args)
	scheduler := services.NewConsecutiveScheduler(downloader)
	workflow := services.NewCourseraWorkflow(scheduler, args, className)
	completed, err := workflow.DownloadModules(modules)
	return completed, err
}
