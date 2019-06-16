package services

import "log"

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

type ConsecutiveScheduler struct {
	AbstractScheduler
}

func NewConsecutiveScheduler(d IDownloader) *ConsecutiveScheduler {
	cs := &ConsecutiveScheduler{}
	as := AbstractScheduler{cs, d}
	cs.AbstractScheduler = as
	return cs
}

func (cs *ConsecutiveScheduler) Download(url string, file string) (string, error) {
	res, err := cs.downloadWrapper(url, file)
	return res, err
}

func (cs *ConsecutiveScheduler) Join(url string) error {
	log.Println("Join is not applicable to ConsecutiveScheduler")
	return nil
}

type ParallelScheduler struct {
	AbstractScheduler
}
