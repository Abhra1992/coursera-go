package services

import "coursera/api"

// WgetDownloader uses the wget tool to download files
type WgetDownloader struct {
	ExternalDownloader
}

// NewWgetDownloader creates a new wget download session
func NewWgetDownloader(session *api.CourseraSession) *WgetDownloader {
	const binary = "wget"
	c := &WgetDownloader{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *WgetDownloader) createCommand(url string, file string) []string {
	return []string{
		url, "-O", file, "--no-cookies", "--no-check-certificate",
	}
}

func (d *WgetDownloader) enableResume(command []string) []string {
	return append(command, "-c")
}

func (d *WgetDownloader) addCookies(command []string, cookies string) []string {
	return append(command, "--header", "Cookie: "+cookies)
}
