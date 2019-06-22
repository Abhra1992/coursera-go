package services

import (
	"coursera/api"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type ExternalDownloader struct {
	IDownloader
	Session *api.CourseraSession
	Binary  string
}

func (ed *ExternalDownloader) startDownload(url string, file string, resume bool) error {
	command := ed.createCommand(url, file)
	command = ed.prepareCookies(command, url)
	if resume {
		command = ed.enableResume(command)
	}
	log.Printf("Executing %s %s", ed.Binary, command)
	process := exec.Command(ed.Binary, command...)
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr
	err := process.Run()
	if err != nil {
		log.Panic("Download Process Failed")
	}
	return nil
}

func (ed *ExternalDownloader) Download(url string, file string, resume bool) error {
	return ed.startDownload(url, file, resume)
}

func (ed *ExternalDownloader) prepareCookies(command []string, url string) []string {
	cookies := ed.Session.Session.RequestOptions.Cookies
	if len(cookies) > 0 {
		command = ed.addCookies(command, getCookieHeader(cookies))
	}
	return command
}

func getCookieHeader(cookies []*http.Cookie) string {
	cookieValues := make([]string, len(cookies))
	for i, c := range cookies {
		cookieValues[i] = fmt.Sprintf("%s=%s", c.Name, c.Value)
	}
	return strings.Join(cookieValues, "; ")
}
