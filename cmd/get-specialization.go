package cmd

import (
	"fmt"
	"sensei/api"
	"sensei/types"
	"sensei/views"
)

// GetSpecialization fetches the details of a specialization
func GetSpecialization(cs *api.Session, name string) (*types.Specialization, error) {
	url := fmt.Sprintf(api.SpecializationURL, name)
	var sr views.SpecializationResponse
	err := cs.GetJSON(url, &sr)
	if err != nil {
		return nil, err
	}
	courses := make([]types.Course, 0, len(sr.Linked.Courses))
	for _, c := range sr.Linked.Courses {
		courses = append(courses, *c.ToModel())
	}
	spz := &types.Specialization{
		Name:    sr.Elements[0].Name,
		Courses: courses,
	}
	return spz, nil
}
