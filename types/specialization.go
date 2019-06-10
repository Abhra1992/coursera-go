package types

type Specialization struct {
	Name    string
	Courses []Course
}

type SpecializationResponse struct {
	Elements []struct {
		CourseIds []string `json:"courseIds"`
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		Slug      string   `json:"slug"`
	} `json:"elements"`
	Linked struct {
		Courses []Course `json:"courses.v1"`
	} `json:"linked"`
}
