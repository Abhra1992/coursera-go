package coursera

import (
	"fmt"
	"sensei/api"
	"sensei/services"
	"sensei/types"
	"strings"

	"github.com/fatih/color"
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
func (e *Extractor) ListCourses() ([]types.CourseResponse, error) {
	course := NewOnDemand(e.Session, "", e.args)
	return course.ListCourses()
}

// GetCourse get the syllabus for a given class
func (e *Extractor) GetCourse(className string) (*types.Course, error) {
	syl, err := e.getOnDemandSyllabus(className)
	if err != nil {
		return nil, err
	}
	course, err := e.parseOnDemandSyllabus(className, syl)
	if err != nil {
		return nil, err
	}
	return course, nil
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

func (e *Extractor) parseOnDemandSyllabus(className string, cm *types.CourseMaterialsResponse) (*types.Course, error) {
	if len(cm.Elements) == 0 {
		return nil, nil
	}
	classID := cm.Elements[0].ID
	color.Green("Syllabus for Course %s", classID)
	od := NewOnDemand(e.Session, classID, e.args)
	var modules []*types.Module
	allModules := cm.GetModuleCollection()
	for _, mr := range allModules {
		color.Yellow("Module [%s] [%s]", mr.ID, mr.Name)
		module, err := e.fillModuleSections(mr, cm, od)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	course := &types.Course{ID: classID, Name: className, Symbol: className, Modules: modules}
	return course, nil
}

func (e *Extractor) fillSectionItems(sr *types.SectionResponse, cm *types.CourseMaterialsResponse,
	course *OnDemand) (*types.Section, error) {
	section := sr.ToModel()
	var items []*types.Item
	allItems := cm.GetItemCollection()
	for _, iid := range sr.ItemIds {
		ir := allItems[iid]
		color.Green("\t\t%s Item [%s] [%s]", strings.Title(ir.ContentSummary.TypeName), ir.ID, ir.Name)
		if ir.IsLocked {
			color.Blue("\t\t\t[Locked] Reason: %s", ir.ItemLockSummary.LockState.ReasonCode)
			continue
		}
		item, err := e.fillItemLinks(ir, course)
		if err != nil {
			return nil, err
		}
		if item.Resources != nil {
			for _, res := range item.Resources {
				maxlen := services.Min(80, len(res.Link))
				color.Cyan("\t\t\t [%s] %s...", res.Extension, res.Link[:maxlen])
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
		color.Magenta("\tSection [%s] [%s]", sr.ID, sr.Name)
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
		color.Red("Unsupported type %s in Item %s %s", item.Type, item.Name, item.ID)
	}
	item.Resources = resx
	return item, nil
}
