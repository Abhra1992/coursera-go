package views

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
