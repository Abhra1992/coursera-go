package views

// Anchor model for supplement assets
type Anchor struct {
	ID   string `json:"id" goquery:"a,text"`
	Link string `json:"url" goquery:"a,[href]"`
}

// AnchorCollection API response for supplement assets
type AnchorCollection struct {
	Elements []*Anchor `json:"elements"`
}

// CoContentAsset API Reponse fragment for asset DTD
type CoContentAsset struct {
	ID        string `goquery:"asset,[id]"`
	Name      string `goquery:"asset,[name]"`
	Extension string `goquery:"asset,[extension]"`
}

// CoContents model for asset html fragment
type CoContents struct {
	Assets  []CoContentAsset `goquery:"co-content"`
	Anchors []Anchor         `goquery:"co-content"`
}
