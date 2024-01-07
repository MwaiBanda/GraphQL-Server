package graph

import (
	"strconv"

	"github.com/MwaiBanda/GraphQL-Server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Todos []*model.Todo
}

func (r *Resolver) AddDefaultTodos() {
	r.Todos = append(r.Todos, &model.Todo{
		ID:   "1",
		Text: "Hello",
		Done: false,
		User: &model.User{
			ID:   "1",
			Name: "Mwai",
		},
	})
}

func (r *Resolver) GetTodos() ([]*model.Todo, error) {
	
	return r.Todos, nil
}

func (r *Resolver) CreateTodos(input model.NewTodo) (*model.Todo, error) {
	lastID, _ := strconv.Atoi(r.Todos[len(r.Todos)-1].ID)
	nextID := strconv.Itoa(lastID + 1)
	todo := &model.Todo{
		ID:   nextID,
		Text: input.Text,
		User: &model.User{
			ID:   input.UserID,
			Name: "user " + input.UserID,
		},
	}
	r.Todos = append(r.Todos, todo)
	return todo, nil
}
