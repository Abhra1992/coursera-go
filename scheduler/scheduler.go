package scheduler

import (
	"coursera/downloader"
	"coursera/types"
)

// IScheduler represents the interface for a download task scheduler
type IScheduler interface {
	Schedule(task DownloadTask)
	Complete()
}

// Scheduler represents an absctract download scheduler
type Scheduler struct {
	IScheduler
	downloader downloader.IDownloader
	args       *types.Arguments
}

func (as *Scheduler) schedule(url string, file string) (string, error) {
	err := as.downloader.Download(url, file, as.args.Resume)
	return url, err
}

// Create instantiates a default scheduler for the downloader
func Create(fd downloader.IDownloader, args *types.Arguments) IScheduler {
	return NewConsecutiveScheduler(fd, args)
}
