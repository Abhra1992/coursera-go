package views

import "sensei/types"

// CourseResponse API reponse for a course
type CourseResponse struct {
	Type string `json:"courseType"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ToModel converts response to model
func (cr *CourseResponse) ToModel() *types.Course {
	return &types.Course{
		ID: cr.ID, Name: cr.Name, Symbol: cr.Slug, Modules: nil,
	}
}
