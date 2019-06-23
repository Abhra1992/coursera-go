package services

// ParallelScheduler schedules download of files in sequence
type ParallelScheduler struct {
	AbstractScheduler
}

// NewParallelScheduler constructor
func NewParallelScheduler(d IDownloader) *ParallelScheduler {
	cs := &ParallelScheduler{}
	as := AbstractScheduler{cs, d}
	cs.AbstractScheduler = as
	return cs
}

// Download starts the download of a url into a file
func (cs *ParallelScheduler) Download(url string, file string) (string, error) {
	panic("Not Implemented")
}

// Join waits for multiple downloads to finish
func (cs *ParallelScheduler) Join() error {
	panic("Not Implemented")
}
