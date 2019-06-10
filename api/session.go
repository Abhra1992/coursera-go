package api

import (
	"coursera/types"
	"fmt"

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

func (cs *CourseraSession) GetSpecialization(name string) (*types.Specialization, error) {
	url := fmt.Sprintf(SpecializationURL, name)
	var sr types.SpecializationResponse
	err := cs.GetJSON(url, &sr)
	if err != nil {
		return nil, err
	}
	spz := &types.Specialization{
		Name:    sr.Elements[0].Name,
		Courses: sr.Linked.Courses,
	}
	return spz, nil
}

func (cs *CourseraSession) getReply(url string,
	post bool, data string, headers map[string]string, quiet bool) (*grequests.Response, error) {
	if headers == nil {
		headers = make(map[string]string)
	}
	res, _ := cs.Session.Get(url, nil)
	if res.Ok != true {
		return res, res.Error
	}
	return res, nil
}

func (cs *CourseraSession) getPage(url string,
	json bool, post bool, data string, headers map[string]string,
	quiet bool, args ...interface{}) (string, error) {
	rurl := fmt.Sprintf(url, args...)
	res, err := cs.getReply(rurl, post, data, headers, quiet)
	if err != nil {
		return "Error", err
	}
	defer res.Close()
	if json {
		body := res.String()
		return body, nil
	} else {
		body := res.String()
		return body, nil
	}
}

func (cs *CourseraSession) getPageJSON(url string,
	post bool, data string, headers map[string]string,
	quiet bool, object interface{}, args ...interface{}) error {
	rurl := fmt.Sprintf(url, args...)
	res, err := cs.getReply(rurl, post, data, headers, quiet)
	if err != nil {
		return err
	}
	defer res.Close()
	err = res.JSON(&object)
	return err
}
