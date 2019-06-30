package types

// Asset model for supplement assets
type Asset struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// AssetResponse API response for supplement assets
type AssetResponse struct {
	Elements []*Asset `json:"elements"`
}

// AssetDefinition API Reponse fragment for asset DTD
type AssetDefinition struct {
	ID        string `goquery:"asset,[id]"`
	Name      string `goquery:"asset,[name]"`
	Extension string `goquery:"asset,[extension]"`
}

// Anchor model for supplement anchors
type Anchor struct {
	Text string `goquery:"a,text"`
	Href string `goquery:"a,[href]"`
}

// AssetPage model for asset html fragment
type AssetPage struct {
	Assets  []AssetDefinition `goquery:"co-content"`
	Anchors []Anchor          `goquery:"co-content"`
}
