package downloader

import (
	"sensei/api"
	"sensei/types"
)

// IDownloader represents the interface of a resource downloader
type IDownloader interface {
	Download(url string, file string, resume bool) error
	startDownload(url string, file string, resume bool) error
	createCommand(url string, file string) []string
	enableResume(command []string) []string
	addCookies(command []string, cookies string) []string
}

// Create instantiates a default downloader for the session
func Create(session *api.Session, args *types.Arguments) IDownloader {
	return NewCurl(session)
}
