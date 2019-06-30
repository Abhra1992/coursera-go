package types

// SupplementsResponse API reponse for an item supplement
type SupplementsResponse struct {
	Elements []struct {
		CourseID string `json:"courseId"`
		ID       string `json:"id"`
		ItemID   string `json:"itemId"`
	} `json:"elements"`
	Linked struct {
		Assets []struct {
			CourseID   string `json:"courseId"`
			Definition struct {
				DtdID string `json:"dtdId"`
				Value string `json:"value"`
			} `json:"definition"`
			ID       string `json:"id"`
			ItemID   string `json:"itemId"`
			TypeName string `json:"typeName"`
		} `json:"openCourseAssets.v1"`
	} `json:"linked"`
}

// LectureAssetsResponse API response for a lecture asset
type LectureAssetsResponse struct {
	Elements []struct {
		CourseID string `json:"courseId"`
		ID       string `json:"id"`
		ItemID   string `json:"itemId"`
	} `json:"elements"`
	Linked struct {
		Assets []struct {
			CourseID   string `json:"courseId"`
			Definition struct {
				DtdID string `json:"dtdId"`
				Value string `json:"value"`
			} `json:"definition"`
			ID       string `json:"id"`
			ItemID   string `json:"itemId"`
			TypeName string `json:"typeName"`
		} `json:"openCourseAssets.v1"`
	} `json:"linked"`
}

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
