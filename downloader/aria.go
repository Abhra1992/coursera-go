package downloader

import "sensei/api"

// Aria2 uses the aria2c tool to download files
type Aria2 struct {
	ExternalDownloader
}

// NewAria2 creates a new aria2c download session
func NewAria2(session *api.Session) *Aria2 {
	const binary = "aria2c"
	c := &Aria2{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *Aria2) createCommand(url string, file string) []string {
	return []string{
		url, "-o", file, "--check-certificate=false", "--log-level=notice",
		"--max-connection-per-server=4", "--min-split-size=1M",
	}
}

func (d *Aria2) enableResume(command []string) []string {
	return append(command, "-c")
}

func (d *Aria2) addCookies(command []string, cookies string) []string {
	return append(command, "--header", "Cookie: "+cookies)
}
