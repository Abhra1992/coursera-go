package types

type Item struct {
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
