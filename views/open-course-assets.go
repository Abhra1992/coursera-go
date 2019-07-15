package views

// OpenCourseAssetsResponse API Response for Open Course assets
type OpenCourseAssetsResponse struct {
	Elements []struct {
		TypeName   string `json:"typeName"`
		Definition struct {
			AssetID string `json:"assetId"`
			Name    string `json:"name"`
			URL     string `json:"string"`
		} `json:"definition"`
		ID string `json:"id"`
	} `json:"elements"`
}
