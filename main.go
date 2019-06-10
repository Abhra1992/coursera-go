package main

import (
	"coursera/cmd"
	"coursera/types"

	"github.com/alexflint/go-arg"
)

func main() {
	var args types.Arguments
	arg.MustParse(&args)

	switch args.ClassType {
	case "spz":
		cmd.HandleSpecialization(args.ClassNames[0])
	case "course":
		cmd.HandleCourses(&args)
	case "list":
		cmd.ListCourses()
	}
}
