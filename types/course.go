package types

// Course model for a course
type Course struct {
	Type   string `json:"courseType"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"slug"`
}
