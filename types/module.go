package types

type Module struct {
	ID       string
	Name     string
	Symbol   string
	Sections []*Section
}

type ModuleResponse struct {
	ID         string   `json:"id"`
	Objectives []string `json:"learningObjectives"`
	LessonIds  []string `json:"lessonIds"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
}

func (mr *ModuleResponse) ToModel() *Module {
	return &Module{
		mr.ID, mr.Name, mr.Slug, nil,
	}
}
