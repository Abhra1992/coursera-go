package views

type openCourseAssetDefinition struct {
	AssetID string `json:"assetId"`
	Name    string `json:"name"`
	URL     string `json:"string"`
}

type openCourseAssetElement struct {
	TypeName   string                    `json:"typeName"`
	Definition openCourseAssetDefinition `json:"definition"`
	ID         string                    `json:"id"`
}

// OpenCourseAssetsResponse API Response for Open Course assets
type OpenCourseAssetsResponse struct {
	Elements []openCourseAssetElement `json:"elements"`
}
