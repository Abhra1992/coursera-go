package types

// Course model for a course
type Course struct {
	ID      string
	Name    string
	Symbol  string
	Modules []*Module
}

// CourseResponse API reponse for a course
type CourseResponse struct {
	Type string `json:"courseType"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ToModel converts response to model
func (cr *CourseResponse) ToModel() *Course {
	return &Course{
		cr.ID, cr.Name, cr.Slug, nil,
	}
}
