package services

import (
	"coursera/types"
	"log"
)

type IWorkflow interface {
	DownloadModules(modules []*types.Module) (bool, error)
}

type CourseraWorkflow struct {
	scheduler IDownloadScheduler
	args      *types.Arguments
	className string
}

func NewCourseraWorkflow(dw IDownloadScheduler, args *types.Arguments, className string) *CourseraWorkflow {
	return &CourseraWorkflow{dw, args, className}
}

func (cw *CourseraWorkflow) DownloadModules(modules []*types.Module) (bool, error) {
	for _, module := range modules {
		log.Printf("MODULE %s", module.Name)
		// for _, section := range module.Sections {
		// 	log.Printf("\tSECTION %s", section.Name)
		// 	for _, item := range section.Items {
		// 		log.Printf("\t\t%s ITEM %s", item.Type, item.Name)
		// 	}
		// }
	}
	res, err := cw.scheduler.Download("url string")
	return len(res) > 0, err
}

func (cw *CourseraWorkflow) HandleResource() {

}
