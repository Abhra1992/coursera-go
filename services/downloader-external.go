package services

import (
	"coursera/api"
	"log"
)

// ExternalDownloader represents an abstract downloader using external tools
type ExternalDownloader struct {
	IDownloader
	Session *api.CourseraSession
	Binary  string
}

// Download downloades a url resource into a file, supports resume
func (ed *ExternalDownloader) Download(url string, file string, resume bool) error {
	return ed.startDownload(url, file, resume)
}

func (ed *ExternalDownloader) startDownload(url string, file string, resume bool) error {
	command := ed.createCommand(url, file)
	command = ed.prepareCookies(command, url)
	if resume {
		command = ed.enableResume(command)
	}
	log.Printf("\t\t> Downloading [%s...] => [%s]", url[:80], file)
	// log.Printf("Executing %s %s", ed.Binary, command)
	// process := exec.Command(ed.Binary, command...)
	// process.Stdout = os.Stdout
	// process.Stderr = os.Stderr
	// err := process.Run()
	// if err != nil {
	// 	log.Panic("Download Process Failed")
	// }
	return nil
}

func (ed *ExternalDownloader) prepareCookies(command []string, url string) []string {
	cookies := ed.Session.Session.RequestOptions.Cookies
	if len(cookies) > 0 {
		command = ed.addCookies(command, api.BuildCookieHeader(cookies))
	}
	return command
}
