package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"sensei/api"
	"sensei/coursera"
	"sensei/scheduler"
	"sensei/services"
	"sensei/types"

	"github.com/fatih/color"
)

// DownloadOnDemandClass downloads a single Coursera On Demand class
func DownloadOnDemandClass(cs *api.Session, className string, args *types.Arguments) (bool, error) {
	var course *types.Course
	// Check if syllabus is cached - if yes, use it
	sf := path.Join(args.Path, fmt.Sprintf("%s-syllabus.json", className))
	if args.CacheSyllabus {
		syllabusExists, err := services.FileExists(sf)
		if err != nil {
			color.Red("Error when checking for existence of syllabus for %s", className)
			return false, err
		}
		if syllabusExists {
			syl, err := ioutil.ReadFile(sf)
			if err != nil {
				color.Red("Could not read cached syllabus for %s", className)
				return false, err
			}
			json.Unmarshal(syl, &course)
		}
	}
	// If no cached syllabus is found, generate the syllabus
	if course == nil {
		// TODO: this extractor step can go inside workflow
		ce := coursera.NewExtractor(cs, args)
		syl, err := ce.ExtractCourse(className)
		if err != nil || syl == nil {
			color.Red("Could not get course syllabus for %s", className)
			return false, err
		}
		course = syl
	}
	// Check if syllabus should be cached - if yes, save it
	if args.CacheSyllabus {
		jsyl, err := json.MarshalIndent(course, "", "\t")
		if err != nil {
			color.Red("Could not cache syllabus for %s", className)
			return false, err
		}
		ioutil.WriteFile(sf, jsyl, 0644)
	}
	if args.OnlySyllabus {
		return true, nil
	}
	ts := scheduler.Create(cs, args)
	workflow := coursera.NewWorkflow(ts, args, className)
	completed, err := workflow.DownloadCourse(course)
	return completed, err
}
