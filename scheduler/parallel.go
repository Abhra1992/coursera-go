package scheduler

import (
	"sensei/downloader"
	"sensei/types"
	"sync"
)

// ParallelScheduler schedules download of files in sequence
type ParallelScheduler struct {
	Scheduler
	pool *sync.WaitGroup
}

// NewParallelScheduler constructor
func NewParallelScheduler(d downloader.IDownloader, args *types.Arguments) *ParallelScheduler {
	cs := &ParallelScheduler{}
	as := Scheduler{cs, d, args}
	cs.Scheduler = as
	cs.pool = &sync.WaitGroup{}
	return cs
}

// Schedule starts the download of a url into a file
func (cs *ParallelScheduler) Schedule(task DownloadTask) {
	cs.pool.Add(1)
	go func() {
		defer cs.pool.Done()
		res, err := cs.schedule(task.URL, task.File)
		task.Callback(res, err)
	}()
}

// Complete waits for multiple downloads to finish
func (cs *ParallelScheduler) Complete() {
	cs.pool.Wait()
}
