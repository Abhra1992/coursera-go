package types

type MembershipsResponse struct {
	Elements []struct {
		CourseID string `json:"courseId"`
		ID       string `json:"id"`
		Role     string `json:"role"`
		UserID   int    `json:"userId"`
	} `json:"elements"`
	Linked struct {
		Courses []Course `json:"courses.v1"`
	} `json:"linked"`
}
