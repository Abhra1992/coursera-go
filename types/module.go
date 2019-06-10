package types

type Module struct {
	ID         string   `json:"id"`
	Objectives []string `json:"learningObjectives"`
	LessonIds  []string `json:"lessonIds"`
	Name       string   `json:"name"`
	Slug       string   `json:"slug"`
}
