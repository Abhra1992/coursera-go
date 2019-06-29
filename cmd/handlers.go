package cmd

import (
	"sensei/api"
	"sensei/types"
	"fmt"
	"log"
)

// HandleSpecialization handles subcommand for specialization
func HandleSpecialization(name string) {
	fmt.Println("Specializations")
	session := api.NewSession(api.CookieFile)
	sp, _ := GetSpecialization(session, name)
	fmt.Println(sp.Name)
	fmt.Println(sp.Courses)
}

// HandleCourses handles subcommand for courses
func HandleCourses(args *types.Arguments) {
	courseNames := args.ClassNames
	log.Printf("Class Names: %s", courseNames)
	session := api.NewSession(api.CookieFile)
	DownloadOnDemandClass(session, courseNames[0], args)
}
