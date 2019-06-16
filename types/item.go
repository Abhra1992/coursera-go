package types

import "strings"

type Item struct {
	ID        string
	Name      string
	Symbol    string
	SectionID string
	ModuleID  string
	Type      string
	Links     map[string]string
}

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

func (ir *ItemResponse) ToModel() *Item {
	return &Item{
		ir.ID, ir.Name, ir.Slug, ir.LessonID, ir.ModuleID, strings.Title(ir.ContentSummary.TypeName), nil,
	}
}
