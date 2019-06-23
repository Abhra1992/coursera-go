package types

// CourseMaterialsResponse API response for course materials
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

// GetItemCollection items in the course materials
func (cm *CourseMaterialsResponse) GetItemCollection() map[string]*ItemResponse {
	imap := make(map[string]*ItemResponse)
	size := len(cm.Linked.Items)
	for i := 0; i < size; i++ {
		it := &cm.Linked.Items[i]
		imap[it.ID] = it
	}
	return imap
}

// GetSectionCollection sections in the course materials
func (cm *CourseMaterialsResponse) GetSectionCollection() map[string]*SectionResponse {
	smap := make(map[string]*SectionResponse)
	size := len(cm.Linked.Lessons)
	for i := 0; i < size; i++ {
		s := &cm.Linked.Lessons[i]
		smap[s.ID] = s
	}
	return smap
}

// GetModuleCollection modules in the course materials
func (cm *CourseMaterialsResponse) GetModuleCollection() map[string]*ModuleResponse {
	mmap := make(map[string]*ModuleResponse)
	size := len(cm.Linked.Modules)
	for i := 0; i < size; i++ {
		m := &cm.Linked.Modules[i]
		mmap[m.ID] = m
	}
	return mmap
}
