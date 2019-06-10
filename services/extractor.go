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

func (e *CourseraExtractor) GetModules(className string, subtitle string) (*types.CourseMaterialsResponse, error) {
	return e.getOnDemandSyllabusJSON(className)
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
