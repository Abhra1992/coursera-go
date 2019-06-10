package services

import (
	"coursera/api"
	"coursera/types"
	"fmt"
	"log"
)

type Extractor interface {
	GetModules() []string
}

type CourseraExtractor struct {
	Session *api.CourseraSession
}

func NewCourseraExtractor(session *api.CourseraSession) *CourseraExtractor {
	return &CourseraExtractor{Session: session}
}

func (e *CourseraExtractor) ListCourses() ([]types.Course, error) {
	course := NewCourseraOnDemand(e.Session)
	return course.ListCourses()
}

func (e *CourseraExtractor) GetModules(className string, subtitle string) ([]*types.Module, error) {
	syl, err := e.getOnDemandSyllabusJSON(className)
	if err != nil {
		return nil, err
	}
	modules, err := e.parseOnDemandSyllabus(className, syl)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

func (e *CourseraExtractor) getOnDemandSyllabus(className string) (string, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, className)
	syl, err := e.Session.GetString(url)
	if err != nil {
		return "", err
	}
	log.Printf("Downloaded %s (%d bytes)", url, len(syl))
	return syl, nil
}

func (e *CourseraExtractor) getOnDemandSyllabusJSON(className string) (*types.CourseMaterialsResponse, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, className)
	var cmr types.CourseMaterialsResponse
	err := e.Session.GetJSON(url, &cmr)
	if err != nil {
		return nil, err
	}
	return &cmr, nil
}

func (e *CourseraExtractor) parseOnDemandSyllabus(className string, cm *types.CourseMaterialsResponse) ([]*types.Module, error) {
	classId := cm.Elements[0].ID
	log.Printf("Parsing syllabus course id %s", classId)
	var modules []*types.Module
	allModules, allSections, allItems := cm.Linked.Modules, cm.Linked.Lessons, cm.Linked.Items
	for m, mr := range allModules {
		log.Printf("Processing Module %d. %s", m, mr.Name)
		module := mr.ToModel()
		// var lessions []types.Section
		log.Println(len(allSections), len(allItems))
		modules = append(modules, module)
	}
	return modules, nil
}
