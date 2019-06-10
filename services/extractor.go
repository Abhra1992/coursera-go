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

func (e *CourseraExtractor) GetModules(cname string, subtitle string) (*types.CourseMaterialsResponse, error) {
	return e.getOnDemandSyllabusJSON(cname)
}

func (e *CourseraExtractor) getOnDemandSyllabus(cname string) (string, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, cname)
	syl, err := e.Session.GetString(url)
	if err != nil {
		return "", err
	}
	log.Printf("Downloaded %s (%d bytes)", url, len(syl))
	return syl, nil
}

func (e *CourseraExtractor) getOnDemandSyllabusJSON(cname string) (*types.CourseMaterialsResponse, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, cname)
	var cmr types.CourseMaterialsResponse
	err := e.Session.GetJSON(url, &cmr)
	if err != nil {
		return nil, err
	}
	return &cmr, nil
}
