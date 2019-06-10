package api

import (
	"github.com/levigross/grequests"
)

type CourseraSession struct {
	Session *grequests.Session
}

func NewCourseraSession(file string) *CourseraSession {
	cj := NewCookieJar(file)
	ro := &grequests.RequestOptions{
		UseCookieJar: true,
		CookieJar:    cj,
	}
	session := grequests.NewSession(ro)
	return &CourseraSession{
		Session: session,
	}
}

// GetString get a String response
func (cs *CourseraSession) GetString(url string) (string, error) {
	res, _ := cs.Session.Get(url, nil)
	if res.Ok != true {
		return "", res.Error
	}
	defer res.Close()
	return res.String(), nil
}

// GetJSON get a JSON response
func (cs *CourseraSession) GetJSON(url string, v interface{}) error {
	res, _ := cs.Session.Get(url, nil)
	if res.Ok != true {
		return res.Error
	}
	defer res.Close()
	return res.JSON(&v)
}
