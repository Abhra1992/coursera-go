package types

// Module model for a course module
type Module struct {
	ID       string
	Name     string
	Symbol   string
	Sections []*Section
}

// ModuleResponse API reponse for a course model
type ModuleResponse struct {
	ID         string   `json:"id"`
	Objectives []string `json:"learningObjectives"`
	LessonIds  []string `json:"lessonIds"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
}

// ToModel converts response to model
func (mr *ModuleResponse) ToModel() *Module {
	return &Module{
		mr.ID, mr.Name, mr.Slug, nil,
	}
}
