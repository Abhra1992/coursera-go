package cmd

import (
	"coursera/api"
	"coursera/types"
	"fmt"
)

func HandleSpecialization(name string) {
	fmt.Println("Specializations")
	session := api.NewCourseraSession(api.CookieFile)
	sp, _ := GetSpecialization(session, name)
	fmt.Println(sp.Name)
	fmt.Println(sp.Courses)
}

func HandleCourses(args *types.Arguments) {
	courseNames := args.ClassNames
	fmt.Println("Downloading Courses")
	fmt.Println(courseNames)
	session := api.NewCourseraSession(api.CookieFile)
	DownloadOnDemandClass(session, courseNames[0], args)
}
