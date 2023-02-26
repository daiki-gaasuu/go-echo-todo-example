package domain

type Todo struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
	// UserId  string
}

func NewTodo(content string) *Todo {
	return &Todo{
		Content: content,
	}
}
