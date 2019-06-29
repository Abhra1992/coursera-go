package cmd

import (
	"coursera/api"
	"coursera/services"
	"coursera/types"
	"fmt"
	"log"
)

// ListCourses lists the courses in which the user is enrolled
func ListCourses(args *types.Arguments) {
	session := api.NewSession(api.CookieFile)
	extractor := services.NewCourseraExtractor(session, args)
	courses, err := extractor.ListCourses()
	if err != nil {
		log.Panicf("Could not list courses")
	}
	for _, c := range courses {
		fmt.Println(c)
	}
}
