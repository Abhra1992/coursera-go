package cmd

import (
	"path"
	"sensei/api"
	"sensei/types"

	"github.com/fatih/color"
)

// HandleSpecialization handles subcommand for specialization
func HandleSpecialization(args *types.Arguments) {
	color.Cyan("Specialization: %s", args.ClassNames)
	cf := path.Join(args.Path, api.CookieFile)
	session := api.NewSession(cf)
	sp, _ := GetSpecialization(session, args.ClassNames[0])
	color.Cyan("Name: %s", sp.Name)
	for _, c := range sp.Courses {
		color.Green("Course Name: %s", c.Name)
		DownloadOnDemandClass(session, c.Slug, args)
	}
}

// HandleCourses handles subcommand for courses
func HandleCourses(args *types.Arguments) {
	courseNames := args.ClassNames
	color.Cyan("Class Names: %s", courseNames)
	cf := path.Join(args.Path, api.CookieFile)
	session := api.NewSession(cf)
	for _, c := range courseNames {
		DownloadOnDemandClass(session, c, args)
	}
}
