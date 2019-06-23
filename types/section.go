package types

// Section model for a module section
type Section struct {
	ID       string
	Name     string
	Symbol   string
	Items    []*Item
	ModuleID string
}

// SectionResponse API response for a module section
type SectionResponse struct {
	ElementIds []string `json:"elementIds"`
	ID         string   `json:"id"`
	ItemIds    []string `json:"itemIds"`
	ModuleID   string   `json:"moduleId"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
	TrackID    string   `json:"trackId"`
}

// ToModel converts response to model
func (sr *SectionResponse) ToModel() *Section {
	return &Section{
		sr.ID, sr.Name, sr.Slug, nil, sr.ModuleID,
	}
}
