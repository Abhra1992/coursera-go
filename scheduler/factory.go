package scheduler

import (
	"sensei/api"
	"sensei/downloader"
	"sensei/types"
)

// Create instantiates a default scheduler for the session
func Create(cs *api.Session, args *types.Arguments) IScheduler {
	fd := downloader.Create(cs, args)
	if args.Jobs > 1 {
		return NewParallelScheduler(fd, args)
	}
	return NewConsecutiveScheduler(fd, args)
}
