package coursera

import (
	"sensei/api"
	"sensei/types"
	"fmt"
	"log"
	"strings"
)

// IExtractor represents the interface for a url extractor from an API
type IExtractor interface {
	GetModules() []string
}

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
func (e *Extractor) ListCourses() ([]types.Course, error) {
	course := NewOnDemand(e.Session, "", e.args)
	return course.ListCourses()
}

// GetModules get the modules for a given class
func (e *Extractor) GetModules(className string) ([]*types.Module, error) {
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

func (e *Extractor) getOnDemandSyllabus(className string) (*types.CourseMaterialsResponse, error) {
	url := fmt.Sprintf(api.CourseMaterialsURL, className)
	var cmr types.CourseMaterialsResponse
	err := e.Session.GetJSON(url, &cmr)
	if err != nil {
		return nil, err
	}
	return &cmr, nil
}

func (e *Extractor) parseOnDemandSyllabus(className string, cm *types.CourseMaterialsResponse) ([]*types.Module, error) {
	classID := cm.Elements[0].ID
	log.Printf("Syllabus for Course %s", classID)
	course := NewOnDemand(e.Session, classID, e.args)
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

func (e *Extractor) fillSectionItems(sr *types.SectionResponse, cm *types.CourseMaterialsResponse,
	course *OnDemand) (*types.Section, error) {
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
		if item.Resources != nil {
			for _, res := range item.Resources {
				log.Printf("\t\t\t [%s] %s...", res.Extension, res.Link[:80])
			}
		}
		items = append(items, item)
	}
	section.Items = items
	return section, nil
}

func (e *Extractor) fillModuleSections(mr *types.ModuleResponse, cm *types.CourseMaterialsResponse,
	course *OnDemand) (*types.Module, error) {
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

func (e *Extractor) fillItemLinks(ir *types.ItemResponse, course *OnDemand) (*types.Item, error) {
	item := ir.ToModel()
	var resx []*types.Resource
	switch item.Type {
	case "Lecture":
		resmap, _ := course.ExtractLinksFromLecture(item.ID)
		resmap.enrich(&resx)
	case "Supplement":
		resmap, _ := course.ExtractLinksFromSupplement(item.ID)
		resmap.enrich(&resx)
	case "PhasedPeer", "GradedProgramming", "UngradedProgramming":
	case "Quiz", "Exam", "Programming", "Notebook":
	default:
		log.Printf("Unsupported type %s in Item %s %s", item.Type, item.Name, item.ID)
	}
	item.Resources = resx
	return item, nil
}
