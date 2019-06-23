package services

// ConsecutiveScheduler schedules download of files in sequence
type ConsecutiveScheduler struct {
	AbstractScheduler
}

// NewConsecutiveScheduler constructor
func NewConsecutiveScheduler(d IDownloader) *ConsecutiveScheduler {
	cs := &ConsecutiveScheduler{}
	as := AbstractScheduler{cs, d}
	cs.AbstractScheduler = as
	return cs
}

// Download starts the download of a url into a file
func (cs *ConsecutiveScheduler) Download(url string, file string) (string, error) {
	res, err := cs.downloadWrapper(url, file)
	return res, err
}

// Join waits for multiple downloads to finish
func (cs *ConsecutiveScheduler) Join() error {
	// Join is not applicable to ConsecutiveScheduler
	return nil
}
