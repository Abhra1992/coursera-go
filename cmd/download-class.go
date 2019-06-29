package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sensei/api"
	"sensei/coursera"
	"sensei/downloader"
	"sensei/scheduler"
	"sensei/services"
	"sensei/types"
)

// DownloadOnDemandClass downloads a single Coursera On Demand class
func DownloadOnDemandClass(cs *api.Session, className string, args *types.Arguments) (bool, error) {
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
		// TODO: this extractor step can go inside workflow
		ce := coursera.NewExtractor(cs, args)
		ems, err := ce.GetModules(className)
		modules = ems
		if err != nil {
			fmt.Println("Error getting Modules")
		}
	}
	// Check if syllabus should be cached - if yes, save it
	if args.CacheSyllabus {
		jsyl, err := json.MarshalIndent(modules, "", "\t")
		if err != nil {
			log.Printf("Could not cache syllabus for %s", className)
			return false, err
		}
		ioutil.WriteFile(sf, jsyl, 0644)
	}
	if args.OnlySyllabus {
		return true, nil
	}
	fd := downloader.Create(cs, args)
	ts := scheduler.Create(fd, args)
	workflow := coursera.NewWorkflow(ts, args, className)
	completed, err := workflow.DownloadModules(modules)
	return completed, err
}
