package types

type CourseMaterialsResponse struct {
	Elements []struct {
		ID        string   `json:"id"`
		ModuleIds []string `json:"moduleIds"`
	} `json:"elements"`
	Linked struct {
		Items   []Item    `json:"onDemandCourseMaterialItems.v2"`
		Lessons []Section `json:"onDemandCourseMaterialLessons.v1"`
		Modules []Module  `json:"onDemandCourseMaterialModules.v1"`
	} `json:"linked"`
}
