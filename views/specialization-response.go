package views

type specializationElement struct {
	CourseIds []string `json:"courseIds"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
}

type specializationLinked struct {
	Courses []CourseResponse `json:"courses.v1"`
}

// SpecializationResponse API response for a specialization
type SpecializationResponse struct {
	Elements []specializationElement `json:"elements"`
	Linked   specializationLinked    `json:"linked"`
}
