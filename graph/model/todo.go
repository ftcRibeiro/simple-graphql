package model

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	UserID      string `json:"user"`
}

type NewTodo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}
