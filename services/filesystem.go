package services

import (
	"net/url"
	"os"
	"runtime"
	"sensei/api"
	"strings"

	"golang.org/x/net/html"
)

// EnsureDirExists makes sure a directory is created before any operations are performed inside it
func EnsureDirExists(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

// FileExists checks if a file is present
func FileExists(fname string) (bool, error) {
	_, err := os.Stat(fname)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var replacer = strings.NewReplacer(":", "-", "/", "-", "<", "-", ">", "-", "\"", "-", "\\", "-", "|", "-", "?", "-", "*", "-", "(", "", ")", "", "\n", " ", "\x00", "-")

// CleanFileName cleans invalid characters from file name
func CleanFileName(fname string) string {
	s := html.UnescapeString(fname)
	q, err := url.QueryUnescape(s)
	if err == nil {
		s = q
	}
	s = replacer.Replace(s)
	s = strings.TrimRight(s, " .")
	return s
}

// CleanURL cleans invalid characters from URL
func CleanURL(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return link
	}
	u = &url.URL{Scheme: u.Scheme, Host: u.Host, Path: u.Path}
	return u.RequestURI()
}

// NormalizeFilePath Prepends device namespace to Windows paths
func NormalizeFilePath(path string) string {
	if runtime.GOOS == "windows" && !strings.HasPrefix(path, api.WindowsUNCPrefix) {
		return api.WindowsUNCPrefix + path
	}
	return path
}
