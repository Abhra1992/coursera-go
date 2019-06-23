package services

import "coursera/api"

// AxelDownloader uses the axel tool to download files
type AxelDownloader struct {
	ExternalDownloader
}

// NewAxelDownloader creates a new axel download session
func NewAxelDownloader(session *api.CourseraSession) *AxelDownloader {
	const binary = "axel"
	c := &AxelDownloader{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *AxelDownloader) createCommand(url string, file string) []string {
	return []string{
		"-o", file, "-n", "4", "-a", url,
	}
}

func (d *AxelDownloader) enableResume(command []string) []string {
	return append(command, "-c")
}

func (d *AxelDownloader) addCookies(command []string, cookies string) []string {
	return append(command, "-H", "Cookie: "+cookies)
}
