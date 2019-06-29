package cmd

import (
	"sensei/api"
	"sensei/coursera"
	"sensei/types"
	"fmt"
	"log"
)

// ListCourses lists the courses in which the user is enrolled
func ListCourses(args *types.Arguments) {
	session := api.NewSession(api.CookieFile)
	extractor := coursera.NewExtractor(session, args)
	courses, err := extractor.ListCourses()
	if err != nil {
		log.Panicf("Could not list courses")
	}
	for _, c := range courses {
		fmt.Println(c)
	}
}
