package cmd

import (
	"fmt"
	"log"
	"path"
	"sensei/api"
	"sensei/coursera"
	"sensei/types"

	"github.com/fatih/color"
)

// DownloadSpecialization handles subcommand for specialization
func DownloadSpecialization(args *types.Arguments) {
	color.Cyan("Specialization: %s", args.ClassNames)
	cf := path.Join(args.Path, api.CookieFile)
	session := api.NewSession(cf)
	sp, _ := GetSpecialization(session, args.ClassNames[0])
	color.Cyan("Name: %s", sp.Name)
	for _, c := range sp.Courses {
		color.Green("Course Name: %s", c.Name)
		DownloadOnDemandClass(session, c.Symbol, args)
	}
}

// DownloadCourses handles subcommand for courses
func DownloadCourses(args *types.Arguments) {
	courseNames := args.ClassNames
	color.Cyan("Class Names: %s", courseNames)
	cf := path.Join(args.Path, api.CookieFile)
	session := api.NewSession(cf)
	for _, c := range courseNames {
		DownloadOnDemandClass(session, c, args)
	}
}

// ListCourses lists the courses in which the user is enrolled
func ListCourses(args *types.Arguments) {
	cf := path.Join(args.Path, api.CookieFile)
	session := api.NewSession(cf)
	extractor := coursera.NewExtractor(session, args)
	courses, err := extractor.ListCourses()
	if err != nil {
		log.Panicf("Could not list courses")
	}
	for _, c := range courses {
		fmt.Println(c)
	}
}
