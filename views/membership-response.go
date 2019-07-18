package views

type membershipElement struct {
	CourseID string `json:"courseId"`
	ID       string `json:"id"`
	Role     string `json:"role"`
	UserID   int    `json:"userId"`
}

type membershipLinked struct {
	Courses []CourseResponse `json:"courses.v1"`
}

// MembershipsResponse API response for course enrolments
type MembershipsResponse struct {
	Elements []membershipElement `json:"elements"`
	Linked   membershipLinked    `json:"linked"`
}
