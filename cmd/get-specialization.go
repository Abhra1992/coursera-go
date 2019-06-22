package cmd

import (
	"coursera/api"
	"coursera/types"
	"fmt"
)

func GetSpecialization(cs *api.CourseraSession, name string) (*types.Specialization, error) {
	url := fmt.Sprintf(api.SpecializationURL, name)
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
