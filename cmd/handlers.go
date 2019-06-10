package cmd

import (
	"coursera/api"
	"coursera/services"
	"coursera/types"
	"fmt"
	"log"
)

func GetSpecialization(cs *api.CourseraSession, name string) (*types.Specialization, error) {
	url := fmt.Sprintf(api.SpecializationURL, name)
	var sr types.SpecializationResponse
	err := cs.GetJSON(url, &sr)
	if err != nil {
		return nil, err
	}
	spz := &types.Specialization{
		Name:    sr.Elements[0].Name,
		Courses: sr.Linked.Courses,
	}
	return spz, nil
}

func HandleSpecialization(name string) {
	fmt.Println("Specializations")
	session := api.NewCourseraSession(api.CookieFile)
	sp, _ := GetSpecialization(session, name)
	fmt.Println(sp.Name)
	fmt.Println(sp.Courses)
}

func DownloadOnDemandClass(cs *api.CourseraSession, className string, args *types.Arguments) (bool, error) {
	extractor := services.NewCourseraExtractor(cs)
	// Check if syllabus is cached - if yes, use it
	modules, err := extractor.GetModules(className, "en")
	if err != nil {
		fmt.Println("Error getting Modules")
	}
	downloader := services.GetDownloader(cs)
	scheduler := services.NewConsecutiveScheduler(downloader)
	workflow := services.NewCourseraWorkflow(scheduler, args, className)
	completed, err := workflow.DownloadModules(modules)
	return completed, err
}

func HandleCourses(args *types.Arguments) {
	courseNames := args.ClassNames
	fmt.Println("Downloading Courses")
	fmt.Println(courseNames)
	session := api.NewCourseraSession(api.CookieFile)
	DownloadOnDemandClass(session, courseNames[0], args)
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
