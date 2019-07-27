package coursera

import (
	"sensei/services"
	"sensei/types"
	"sensei/views"
	"strings"

	"github.com/fatih/color"
)

func (od *OnDemand) buildCourse(className string, cm *views.CourseMaterialsResponse) (*types.Course, error) {
	var modules []*types.Module
	allModules := cm.GetModuleCollection()
	for _, mr := range allModules {
		color.Yellow("Module [%s] [%s]", mr.ID, mr.Name)
		module, err := od.fillModuleSections(mr, cm)
		if err != nil {
			return nil, err
		}
		modules = append(modules, module)
	}
	course := &types.Course{ID: od.classID, Name: className, Symbol: className, Modules: modules}
	return course, nil
}

func (od *OnDemand) fillModuleSections(mr *views.ModuleResponse, cm *views.CourseMaterialsResponse) (*types.Module, error) {
	module := mr.ToModel()
	var sections []*types.Section
	allSections := cm.GetSectionCollection()
	for _, sid := range mr.LessonIds {
		sr := allSections[sid]
		color.Magenta("\tSection [%s] [%s]", sr.ID, sr.Name)
		section, err := od.fillSectionItems(sr, cm)
		if err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}
	module.Sections = sections
	return module, nil
}

func (od *OnDemand) fillSectionItems(sr *views.SectionResponse, cm *views.CourseMaterialsResponse) (*types.Section, error) {
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
		item, err := od.fillItemLinks(ir)
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

func (od *OnDemand) fillItemLinks(ir *views.ItemResponse) (*types.Item, error) {
	item := ir.ToModel()
	var resx []*types.Resource
	switch item.Type {
	case "Lecture":
		resmap, _ := od.extractLinksFromLecture(item.ID)
		resmap.enrich(&resx)
	case "Supplement":
		resmap, _ := od.extractLinksFromSupplement(item.ID)
		resmap.enrich(&resx)
	case "PhasedPeer", "GradedProgramming", "UngradedProgramming":
	case "Quiz", "Exam", "Programming", "Notebook":
	default:
		color.Red("Unsupported type %s in Item %s %s", item.Type, item.Name, item.ID)
	}
	item.Resources = resx
	return item, nil
}
