package dto

type Todo struct {
	ID   *uint   `json:"id"`
	Name *string `json:"name"`
}

type CreatTodo struct {
	Name *string `json:"name"`
}
