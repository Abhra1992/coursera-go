package services

// IDownloadScheduler represents the interface for a download task scheduler
type IDownloadScheduler interface {
	Download(url string, file string) (string, error)
	Join() error
}

// AbstractScheduler represents an absctract download scheduler
type AbstractScheduler struct {
	IDownloadScheduler
	downloader IDownloader
}

func (as *AbstractScheduler) downloadWrapper(url string, file string) (string, error) {
	err := as.downloader.Download(url, file, true)
	return url, err
}
