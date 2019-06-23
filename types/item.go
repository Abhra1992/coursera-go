package types

import "strings"

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

// ItemResponse API response for a section item
type ItemResponse struct {
	ContentSummary struct {
		TypeName string `json:"typeName"`
	} `json:"contentSummary"`
	ItemLockSummary struct {
		LockState struct {
			LockStatus string `json:"lockStatus"`
			ReasonCode string `json:"reasonCode"`
		} `json:"lockState"`
	} `json:"itemLockSummary"`
	ID       string `json:"id"`
	IsLocked bool   `json:"isLocked"`
	LessonID string `json:"lessonId"`
	ModuleID string `json:"moduleId"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	TrackID  string `json:"trackId"`
}

// ToModel converts response to model
func (ir *ItemResponse) ToModel() *Item {
	return &Item{
		ir.ID, ir.Name, ir.Slug, ir.LessonID, ir.ModuleID, strings.Title(ir.ContentSummary.TypeName), nil,
	}
}
