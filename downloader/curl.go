package downloader

import "coursera/api"

// Curl uses the curl tool to download files
type Curl struct {
	ExternalDownloader
}

// NewCurl creates a new curl download session
func NewCurl(session *api.Session) *Curl {
	const binary = "C:/Windows/System32/curl.exe"
	c := &Curl{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *Curl) createCommand(url string, file string) []string {
	return []string{
		url, "-k", "-#", "-L", "-o", file,
	}
}

func (d *Curl) enableResume(command []string) []string {
	return append(command, "-C", "-")
}

func (d *Curl) addCookies(command []string, cookies string) []string {
	return append(command, "--cookie", cookies)
}
