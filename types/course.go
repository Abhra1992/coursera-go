package types

// Course model for a course
type Course struct {
	ID      string
	Name    string
	Symbol  string
	Modules []*Module
}
