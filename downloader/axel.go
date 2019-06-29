package downloader

import "coursera/api"

// Axel uses the axel tool to download files
type Axel struct {
	ExternalDownloader
}

// NewAxel creates a new axel download session
func NewAxel(session *api.Session) *Axel {
	const binary = "axel"
	c := &Axel{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *Axel) createCommand(url string, file string) []string {
	return []string{
		"-o", file, "-n", "4", "-a", url,
	}
}

func (d *Axel) enableResume(command []string) []string {
	return append(command, "-c")
}

func (d *Axel) addCookies(command []string, cookies string) []string {
	return append(command, "-H", "Cookie: "+cookies)
}
