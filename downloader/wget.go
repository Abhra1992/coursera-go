package downloader

import "coursera/api"

// Wget uses the wget tool to download files
type Wget struct {
	ExternalDownloader
}

// NewWget creates a new wget download session
func NewWget(session *api.Session) *Wget {
	const binary = "wget"
	c := &Wget{}
	e := ExternalDownloader{c, session, binary}
	c.ExternalDownloader = e
	return c
}

func (d *Wget) createCommand(url string, file string) []string {
	return []string{
		url, "-O", file, "--no-cookies", "--no-check-certificate",
	}
}

func (d *Wget) enableResume(command []string) []string {
	return append(command, "-c")
}

func (d *Wget) addCookies(command []string, cookies string) []string {
	return append(command, "--header", "Cookie: "+cookies)
}
