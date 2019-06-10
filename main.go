package main

import (
	"coursera/cmd"
	"coursera/types"
	"flag"
)

func main() {
	classType := flag.String("type", "course", "Class Type - Course or Specialization")
	var courseNames types.SArray
	flag.Var(&courseNames, "name", "Course Name")
	flag.Parse()

	switch *classType {
	case "spz":
		cmd.HandleSpecialization(courseNames[0])
	case "course":
		cmd.HandleCourses(courseNames)
	case "list":
		cmd.ListCourses()
	}
}
