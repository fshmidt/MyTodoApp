package service

import (
	mytodo "github.com/fshmidt/todo-app"
	"github.com/fshmidt/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user mytodo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list mytodo.Todolist) (int, error)
	GetAll(userId int) ([]mytodo.Todolist, error)
	GetById(userId, listId int) (mytodo.Todolist, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input mytodo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item mytodo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]mytodo.TodoItem, error)
	GetById(userId, itemId int) (mytodo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input mytodo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
