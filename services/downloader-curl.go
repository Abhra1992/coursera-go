package services

import "coursera/api"

type CurlDownloader struct {
	ExternalDownloader
}

func NewCurlDownloader(session *api.CourseraSession) *CurlDownloader {
	const binary = "C:/Windows/System32/curl.exe"
	c := &CurlDownloader{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *CurlDownloader) createCommand(url string, file string) []string {
	return []string{
		url, "-k", "-#", "-L", "-o", file,
	}
}

func (d *CurlDownloader) enableResume(command []string) []string {
	return append(command, "-C", "-")
}

func (d *CurlDownloader) addCookies(command []string, cookies string) []string {
	return append(command, "--cookie", cookies)
}