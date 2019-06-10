package types

type CourseMaterialsResponse struct {
	Elements []struct {
		ID        string   `json:"id"`
		ModuleIds []string `json:"moduleIds"`
	} `json:"elements"`
	Linked struct {
		Items   []ItemResponse    `json:"onDemandCourseMaterialItems.v2"`
		Lessons []SectionResponse `json:"onDemandCourseMaterialLessons.v1"`
		Modules []ModuleResponse  `json:"onDemandCourseMaterialModules.v1"`
	} `json:"linked"`
}
