package main

import (
	"sensei/cmd"
	"sensei/types"

	"github.com/alexflint/go-arg"
)

func main() {
	var args types.Arguments
	arg.MustParse(&args)

	switch args.ClassType {
	case "spz":
		cmd.HandleSpecialization(&args)
	case "course":
		cmd.HandleCourses(&args)
	case "list":
		cmd.ListCourses(&args)
	}
}
