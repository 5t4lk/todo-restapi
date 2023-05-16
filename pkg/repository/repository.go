package repository

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *DBConnection) *Repository {
	return &Repository{}
}
