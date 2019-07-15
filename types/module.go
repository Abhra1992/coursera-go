package types

// Module model for a course module
type Module struct {
	ID       string
	Name     string
	Symbol   string
	Sections []*Section
}
