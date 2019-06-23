package services

import (
	"coursera/api"
	"coursera/types"
)

// IDownloader represents the interface of a resource downloader
type IDownloader interface {
	Download(url string, file string, resume bool) error
	startDownload(url string, file string, resume bool) error
	createCommand(url string, file string) []string
	enableResume(command []string) []string
	addCookies(command []string, cookies string) []string
}

// GetDownloader instantiates a default downloader for the session
func GetDownloader(session *api.CourseraSession, args *types.Arguments) IDownloader {
	return NewCurlDownloader(session)
}
