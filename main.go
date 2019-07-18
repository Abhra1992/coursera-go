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
		cmd.DownloadSpecialization(&args)
	case "course":
		cmd.DownloadCourses(&args)
	case "list":
		cmd.ListCourses(&args)
	}
}
