package views

import (
	"sensei/types"
	"strings"
)

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
func (ir *ItemResponse) ToModel() *types.Item {
	return &types.Item{
		ID: ir.ID, Name: ir.Name, Symbol: ir.Slug,
		SectionID: ir.LessonID, ModuleID: ir.ModuleID,
		Type: strings.Title(ir.ContentSummary.TypeName), Resources: nil,
	}
}
