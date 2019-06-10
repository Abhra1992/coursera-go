package types

type Section struct {
	ElementIds []string `json:"elementIds"`
	ID         string   `json:"id"`
	ItemIds    []string `json:"itemIds"`
	ModuleID   string   `json:"moduleId"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
	TrackID    string   `json:"trackId"`
}
