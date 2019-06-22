package services

type IDownloadScheduler interface {
	Download(url string, file string) (string, error)
	Join(url string) error
}

type AbstractScheduler struct {
	IDownloadScheduler
	downloader IDownloader
}

func (as *AbstractScheduler) downloadWrapper(url string, file string) (string, error) {
	err := as.downloader.Download(url, file, true)
	return url, err
}
