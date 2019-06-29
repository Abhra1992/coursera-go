package cmd

import (
	"fmt"
	"log"
	"path"
	"sensei/api"
	"sensei/coursera"
	"sensei/types"
)

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
