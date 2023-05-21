package todo

type TodoList struct {
	Id          string `json:"-" bson:"id"`
	Title       string `json:"title" bson:"title" binding:"required"`
	Description string `json:"description" bson:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
