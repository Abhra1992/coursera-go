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
	session := api.NewCourseraSession(api.CookieFile)
	DownloadOnDemandClass(courseNames[0], session, true)
}

func DownloadOnDemandClass(className string, session *api.CourseraSession, cache bool) {
	extractor := services.NewCourseraExtractor(session)
	// Check if syllabus is cached - if yes, use it
	modules, err := extractor.GetModules(className, "en")
	if err != nil {
		fmt.Println("Error getting Modules")
	}
	fmt.Println(len(modules.Elements))
	downloader := services.GetDownloader(session)
	downloader.Download("url string", "file string", false)
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
