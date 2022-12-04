package repository

import (
	mytodo "github.com/fshmidt/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user mytodo.User) (int, error)
	GetUser(username, password string) (mytodo.User, error)
}

type TodoList interface {
	Create(userId int, list mytodo.Todolist) (int, error)
	GetAll(userId int) ([]mytodo.Todolist, error)
	GetById(userId, listId int) (mytodo.Todolist, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input mytodo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item mytodo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]mytodo.TodoItem, error)
	GetById(userId, itemId int) (mytodo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input mytodo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodolistPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
