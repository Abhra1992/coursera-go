package types

type Course struct {
	CourseType string `json:"courseType"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
}
