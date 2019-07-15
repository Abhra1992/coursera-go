package views

import "sensei/types"

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
func (sr *SectionResponse) ToModel() *types.Section {
	return &types.Section{
		ID: sr.ID, Name: sr.Name, Symbol: sr.Slug,
		Items: nil, ModuleID: sr.ModuleID,
	}
}
