package types

// MembershipsResponse API response for course enrolments
type MembershipsResponse struct {
	Elements []struct {
		CourseID string `json:"courseId"`
		ID       string `json:"id"`
		Role     string `json:"role"`
		UserID   int    `json:"userId"`
	} `json:"elements"`
	Linked struct {
		Courses []CourseResponse `json:"courses.v1"`
	} `json:"linked"`
}
