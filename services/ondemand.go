package services

import (
	"coursera/api"
	"coursera/types"
)

type CourseraOnDemand struct {
	Session *api.CourseraSession
}

func NewCourseraOnDemand(session *api.CourseraSession) *CourseraOnDemand {
	return &CourseraOnDemand{Session: session}
}

func (od *CourseraOnDemand) ObtainUserId() (int, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURLLimit1, &mr)
	if err != nil {
		return -1, err
	}
	return mr.Elements[0].UserID, nil
}

func (od *CourseraOnDemand) ListCourses() ([]types.Course, error) {
	var mr types.MembershipsResponse
	err := od.Session.GetJSON(api.MembershipsURL, &mr)
	if err != nil {
		return nil, err
	}
	return mr.Linked.Courses, nil
}
