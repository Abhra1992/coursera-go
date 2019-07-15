package views

import "sensei/types"

// ModuleResponse API reponse for a course model
type ModuleResponse struct {
	ID         string   `json:"id"`
	Objectives []string `json:"learningObjectives"`
	LessonIds  []string `json:"lessonIds"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
}

// ToModel converts response to model
func (mr *ModuleResponse) ToModel() *types.Module {
	return &types.Module{
		ID: mr.ID, Name: mr.Name, Symbol: mr.Slug, Sections: nil,
	}
}
