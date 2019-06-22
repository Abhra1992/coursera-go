package cmd

import (
	"coursera/api"
	"coursera/services"
	"coursera/types"
	"fmt"
	"log"
)

func ListCourses(args *types.Arguments) {
	session := api.NewCourseraSession(api.CookieFile)
	extractor := services.NewCourseraExtractor(session, args)
	courses, err := extractor.ListCourses()
	if err != nil {
		log.Panicf("Could not list courses")
	}
	for _, c := range courses {
		fmt.Println(c)
	}
}
