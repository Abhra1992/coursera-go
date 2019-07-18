package views

type assetElement struct {
	CourseID string `json:"courseId"`
	ID       string `json:"id"`
	ItemID   string `json:"itemId"`
}

type assetDefinition struct {
	DtdID string `json:"dtdId"`
	Value string `json:"value"`
}

// TODO: Inherit from or Embed assetElement
type itemAsset struct {
	CourseID   string          `json:"courseId"`
	Definition assetDefinition `json:"definition"`
	ID         string          `json:"id"`
	ItemID     string          `json:"itemId"`
	TypeName   string          `json:"typeName"`
}

type assetsLinked struct {
	Assets []itemAsset `json:"openCourseAssets.v1"`
}

// AssetsResponse API response for a lecture asset
type AssetsResponse struct {
	Elements []assetElement `json:"elements"`
	Linked   assetsLinked   `json:"linked"`
}
