package types

// Specialization model for a specialization
type Specialization struct {
	Name    string
	Courses []CourseResponse
}

// SpecializationResponse API response for a specialization
type SpecializationResponse struct {
	Elements []struct {
		CourseIds []string `json:"courseIds"`
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		Slug      string   `json:"slug"`
	} `json:"elements"`
	Linked struct {
		Courses []CourseResponse `json:"courses.v1"`
	} `json:"linked"`
}
