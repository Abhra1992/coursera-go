package types

// Asset model for supplement assets
type Asset struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// AssetDefinition API Reponse fragment for asset DTD
type AssetDefinition struct {
	ID        string `goquery:"asset,[id]"`
	Name      string `goquery:"asset,[name]"`
	Extension string `goquery:"asset,[extension]"`
}

// AssetResponse API response for supplement assets
type AssetResponse struct {
	Assets []*Asset `json:"elements"`
}
