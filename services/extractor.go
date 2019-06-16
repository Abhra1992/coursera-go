package services

import (
	"coursera/api"
	"coursera/types"
	"fmt"
	"log"
	"strings"
)

type Extractor interface {
	GetModules() []string
}

type CourseraExtractor struct {
	Session *api.CourseraSession
	args    *types.Arguments
}

func NewCourseraExtractor(session *api.CourseraSession, args *types.Arguments) *CourseraExtractor {
	return &CourseraExtractor{Session: session, args: args}
}

func (e *CourseraExtractor) ListCourses() ([]types.Course, error) {
	course := NewCourseraOnDemand(e.Session, "", e.args)
	return course.ListCourses()
}

func (e *CourseraExtractor) GetModules(className string) ([]*types.Module, error) {
	syl, err := e.getOnDemandSyllabus(className)
	if err != nil {
		return nil, err
	}
	modules, err := e.parseOnDemandSyllabus(className, syl)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

// DEPRECATED
func (e *CourseraExtractor) getOnDemandSyllabusString(className string) (string, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, className)
	syl, err := e.Session.GetString(url)
	if err != nil {
		return "", err
	}
	log.Printf("Downloaded %s (%d bytes)", url, len(syl))
	return syl, nil
}

func (e *CourseraExtractor) getOnDemandSyllabus(className string) (*types.CourseMaterialsResponse, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, className)
	var cmr types.CourseMaterialsResponse
	err := e.Session.GetJSON(url, &cmr)
	if err != nil {
		return nil, err
	}
	return &cmr, nil
}

func (e *CourseraExtractor) parseOnDemandSyllabus(className string, cm *types.CourseMaterialsResponse) ([]*types.Module, error) {
	classID := cm.Elements[0].ID
	log.Printf("Syllabus for Course %s", classID)
	course := NewCourseraOnDemand(e.Session, classID, e.args)
	var modules []*types.Module
	allModules := cm.GetModuleCollection()
	for _, mr := range allModules {
		log.Printf("Module [%s] [%s]", mr.ID, mr.Name)
		module, err := e.fillModuleSections(mr, cm, course)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	return modules, nil
}

func (e *CourseraExtractor) fillSectionItems(sr *types.SectionResponse, cm *types.CourseMaterialsResponse,
	course *CourseraOnDemand) (*types.Section, error) {
	section := sr.ToModel()
	var items []*types.Item
	allItems := cm.GetItemCollection()
	for _, iid := range sr.ItemIds {
		ir := allItems[iid]
		log.Printf("\t\t%s Item [%s] [%s]", strings.Title(ir.ContentSummary.TypeName), ir.ID, ir.Name)
		if ir.IsLocked {
			log.Printf("\t\t\t[Locked] Reason: %s", ir.ItemLockSummary.LockState.ReasonCode)
			continue
		}
		item, err := e.fillItemLinks(ir, course)
		if err != nil {
			return nil, err
		}
		if item.Links != nil {
			for key, link := range item.Links {
				log.Printf("\t\t\t [%s] %s...", key, link[:50])
			}
		}
		items = append(items, item)
	}
	section.Items = items
	return section, nil
}

func (e *CourseraExtractor) fillModuleSections(mr *types.ModuleResponse, cm *types.CourseMaterialsResponse,
	course *CourseraOnDemand) (*types.Module, error) {
	module := mr.ToModel()
	var sections []*types.Section
	allSections := cm.GetSectionCollection()
	for _, sid := range mr.LessonIds {
		sr := allSections[sid]
		log.Printf("\tSection [%s] [%s]", sr.ID, sr.Name)
		section, err := e.fillSectionItems(sr, cm, course)
		if err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}
	module.Sections = sections
	return module, nil
}

func (e *CourseraExtractor) fillItemLinks(ir *types.ItemResponse, course *CourseraOnDemand) (*types.Item, error) {
	item := ir.ToModel()
	var links map[string]string
	switch item.Type {
	case "Lecture":
		links, _ = course.ExtractLinksFromLecture(item.ID)
	case "Supplement":
		links, _ = course.ExtractLinksFromSupplement(item.ID)
	case "PhasedPeer", "GradedProgramming", "UngradedProgramming":
	case "Quiz", "Exam", "Programming", "Notebook":
	default:
		log.Printf("Unsupported type %s in Item %s %s", item.Type, item.Name, item.ID)
	}
	item.Links = links
	return item, nil
}
