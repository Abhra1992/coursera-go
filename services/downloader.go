package services

import (
	"coursera/api"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type IDownloader interface {
	Download(url string, file string, resume bool) error
	startDownload(url string, file string, resume bool) error
	createCommand(url string, file string) []string
	enableResume(command []string) []string
	addCookies(command []string, cookies string) []string
}

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
	// process := exec.Command(ed.Binary, command...)
	// err := process.Run()
	// if err != nil {
	// 	log.Panic("Download Process Failed")
	// }
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

type WgetDownloader struct {
	ExternalDownloader
}

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
		url, "-k", "-#", "=L", "-o", file,
	}
}

func (d *CurlDownloader) enableResume(command []string) []string {
	return append(command, "-C", "-")
}

func (d *CurlDownloader) addCookies(command []string, cookies string) []string {
	return append(command, "--cookie", cookies)
}

func GetDownloader(session *api.CourseraSession) IDownloader {
	return NewCurlDownloader(session)
}
