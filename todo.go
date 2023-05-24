package todo

import "errors"

type TodoList struct {
	Id          string `json:"_id" bson:"_id"`
	Title       string `json:"title" bson:"title" binding:"required"`
	Description string `json:"description" bson:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          string `json:"_id" bson:"_id"`
	Title       string `json:"title" bson:"title" binding:"required"`
	Description string `json:"description" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

func (i *UpdateListInput) Validate() error {
	if i.Title == "" && i.Description == "" {
		return errors.New("request of update structure has no values")
	}
	return nil
}

type UpdateItemInput struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Done        *bool  `json:"done" bson:"done"`
}

func (i *UpdateItemInput) Validate() error {
	if i.Title == "" && i.Description == "" && i.Done == nil {
		return errors.New("request of update structure has no values")
	}
	return nil
}
