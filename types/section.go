package types

// Section model for a module section
type Section struct {
	ID       string
	Name     string
	Symbol   string
	Items    []*Item
	ModuleID string
}
