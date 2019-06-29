package api

import (
	"github.com/levigross/grequests"
)

// Session is a Coursera download session
type Session struct {
	*grequests.Session
}

// NewSession constructor
func NewSession(file string) *Session {
	cj := NewCookieJar(file)
	ro := &grequests.RequestOptions{
		UseCookieJar: true,
		CookieJar:    cj,
	}
	session := grequests.NewSession(ro)
	return &Session{session}
}

// GetString get a String response from an API
func (cs *Session) GetString(url string) (string, error) {
	res, _ := cs.Get(url, nil)
	if res.Ok != true {
		return "", res.Error
	}
	defer res.Close()
	return res.String(), nil
}

// GetJSON get a JSON response from an API
func (cs *Session) GetJSON(url string, v interface{}) error {
	res, _ := cs.Get(url, nil)
	if res.Ok != true {
		return res.Error
	}
	defer res.Close()
	return res.JSON(&v)
}
