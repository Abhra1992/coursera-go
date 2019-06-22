package services

import (
	"coursera/api"
)

type IDownloader interface {
	Download(url string, file string, resume bool) error
	startDownload(url string, file string, resume bool) error
	createCommand(url string, file string) []string
	enableResume(command []string) []string
	addCookies(command []string, cookies string) []string
}

func GetDownloader(session *api.CourseraSession) IDownloader {
	return NewCurlDownloader(session)
}
