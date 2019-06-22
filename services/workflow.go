package services

import (
	"coursera/types"
)

type IWorkflow interface {
	DownloadModules(modules []*types.Module) (bool, error)
}
