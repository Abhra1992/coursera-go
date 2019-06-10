package cmd

import (
	"coursera/api"
	"coursera/services"
	"fmt"
	"log"
)

func HandleSpecialization(name string) {
	fmt.Println("Specializations")
	session := api.NewCourseraSession(api.CookieFile)
	sp, _ := session.GetSpecialization(name)
	fmt.Println(sp.Name)
	fmt.Println(sp.Courses)
}

func HandleCourses(courseNames []string) {
	fmt.Println("Downloading Courses")
	fmt.Println(courseNames)
}

func DownloadOnDemandClass(cname string, session *api.CourseraSession, cache bool) {
	extractor := services.NewCourseraExtractor(session)
	// Check if syllabus is cached - if yes, use it
	extractor.GetModules(cname, "en")
}

func ListCourses() {
	session := api.NewCourseraSession(api.CookieFile)
	extractor := services.NewCourseraExtractor(session)
	courses, err := extractor.ListCourses()
	if err != nil {
		log.Panicf("Could not list courses")
	}
	for _, c := range courses {
		fmt.Println(c)
	}
}
