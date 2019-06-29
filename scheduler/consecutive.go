package scheduler

import (
	"coursera/downloader"
	"coursera/types"
)

// ConsecutiveScheduler schedules download of files in sequence
type ConsecutiveScheduler struct {
	Scheduler
}

// NewConsecutiveScheduler constructor
func NewConsecutiveScheduler(d downloader.IDownloader, args *types.Arguments) *ConsecutiveScheduler {
	cs := &ConsecutiveScheduler{}
	as := Scheduler{cs, d, args}
	cs.Scheduler = as
	return cs
}

// Schedule starts the download of a url into a file
func (cs *ConsecutiveScheduler) Schedule(task DownloadTask) {
	res, err := cs.schedule(task.URL, task.File)
	task.Callback(res, err)
}

// Complete waits for multiple downloads to finish
func (cs *ConsecutiveScheduler) Complete() {
	// Complete is not applicable to ConsecutiveScheduler
}
