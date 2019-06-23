package services

import (
	"coursera/types"
)

// IWorkflow represents the interface for a download workflow
type IWorkflow interface {
	DownloadModules(modules []*types.Module) (bool, error)
}
