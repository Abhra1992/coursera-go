package services

import "coursera/api"

// Aria2Downloader uses the aria2c tool to download files
type Aria2Downloader struct {
	ExternalDownloader
}

// NewAria2Downloader creates a new aria2c download session
func NewAria2Downloader(session *api.CourseraSession) *Aria2Downloader {
	const binary = "aria2c"
	c := &Aria2Downloader{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *Aria2Downloader) createCommand(url string, file string) []string {
	return []string{
		url, "-o", file, "--check-certificate=false", "--log-level=notice",
		"--max-connection-per-server=4", "--min-split-size=1M",
	}
}

func (d *Aria2Downloader) enableResume(command []string) []string {
	return append(command, "-c")
}

func (d *Aria2Downloader) addCookies(command []string, cookies string) []string {
	return append(command, "--header", "Cookie: "+cookies)
}
