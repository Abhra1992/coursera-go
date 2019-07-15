package types

// Item model for a section item
type Item struct {
	ID        string
	Name      string
	Symbol    string
	SectionID string
	ModuleID  string
	Type      string
	Resources []*Resource
}
