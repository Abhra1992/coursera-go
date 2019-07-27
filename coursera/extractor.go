package coursera

import (
	"fmt"
	"sensei/api"
	"sensei/types"
	"sensei/views"

	"github.com/fatih/color"
)

// Extractor extracts links from the Coursera API
type Extractor struct {
	Session *api.Session
	args    *types.Arguments
}

// NewExtractor creates a new Coursera Extractor
func NewExtractor(session *api.Session, args *types.Arguments) *Extractor {
	return &Extractor{Session: session, args: args}
}

// ListCourses list the courses the user has enrolled in
func (e *Extractor) ListCourses() ([]views.CourseResponse, error) {
	course := NewOnDemand(e.Session, "", e.args)
	return course.ListCourses()
}

// ExtractCourse get the syllabus for a given class
func (e *Extractor) ExtractCourse(className string) (*types.Course, error) {
	syl, err := e.getSyllabus(className)
	if err != nil {
		return nil, err
	}
	course, err := e.convertSyllabus(className, syl)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (e *Extractor) getSyllabus(className string) (*views.CourseMaterialsResponse, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, className)
	var cmr views.CourseMaterialsResponse
	if err := e.Session.GetJSON(url, &cmr); err != nil {
		return nil, err
	}
	return &cmr, nil
}

func (e *Extractor) convertSyllabus(className string, cm *views.CourseMaterialsResponse) (*types.Course, error) {
	if len(cm.Elements) == 0 {
		return nil, nil
	}
	classID := cm.Elements[0].ID
	color.Green("Syllabus for Course %s", classID)
	od := NewOnDemand(e.Session, classID, e.args)
	return od.buildCourse(className, cm)
}
