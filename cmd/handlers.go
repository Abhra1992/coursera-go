package cmd

import (
	"fmt"
	"sensei/api"
	"sensei/types"

	"github.com/fatih/color"
)

// HandleSpecialization handles subcommand for specialization
func HandleSpecialization(name string) {
	color.Cyan("Specializations")
	session := api.NewSession(api.CookieFile)
	sp, _ := GetSpecialization(session, name)
	fmt.Println(sp.Name)
	fmt.Println(sp.Courses)
}

// HandleCourses handles subcommand for courses
func HandleCourses(args *types.Arguments) {
	courseNames := args.ClassNames
	color.Cyan("Class Names: %s", courseNames)
	session := api.NewSession(api.CookieFile)
	DownloadOnDemandClass(session, courseNames[0], args)
}
