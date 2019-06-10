package services

import "coursera/types"

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
	// modules
	res, err := cw.scheduler.Download("url string")
	return len(res) > 0, err
}
