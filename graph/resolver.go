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
			Name: input.UserID,
		},
	}
	r.Todos = append(r.Todos, todo)
	return todo, nil
}

func (r *Resolver) UpdateTodos(input model.UpdateTodo) (*model.Todo, error) {
	todo := new(model.Todo)
	for i, todo := range r.Todos {
		if todo.ID == input.ID {
			r.Todos[i].Text = input.Text
			todo = r.Todos[i]
			return todo, nil
		}
	}
	return todo, nil
}

func (r *Resolver) DeleteTodos(input model.DeleteTodo) (*model.Todo, error) {
	updatedArray := make([]*model.Todo, (len(r.Todos) - 1))
	todo := new(model.Todo)
	for i, v := range r.Todos {
    	if r.Todos[i].ID != input.ID {
        	updatedArray[i] = v
    	} else {
        	todo = r.Todos[i]
    	}
	}
	r.Todos = updatedArray
	return todo, nil
}