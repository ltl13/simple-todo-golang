package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	models "local/todo/models"
	repo "local/todo/repositories"
	"net/http"
	"strconv"
)

type TodoController struct {
	repository *repo.TodoRepository
}

func (x *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("This URL doesn't support method %s", r.Method), 405)
		return
	}
	result, statusCode, err := x.repository.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	resp, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func (x *TodoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("This URL doesn't support method %s", r.Method), 405)
		return
	}
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}
	result, statusCode, err := x.repository.GetTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}
	resp, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resp)
}

func (x *TodoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("This URL doesn't support method %s", r.Method), 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request body", 400)
		return
	}
	var todo models.Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		http.Error(w, "Bad request body", 400)
	}
	statusCode, err := x.repository.AddTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
	w.WriteHeader(statusCode)
}

func (x *TodoController) RemoveTodoByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, fmt.Sprintf("This URL doesn't support method %s", r.Method), 405)
		return
	}
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}
	statusCode, err := x.repository.RemoveTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
	w.WriteHeader(statusCode)
}

func (x *TodoController) UpdateTodoByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, fmt.Sprintf("This URL doesn't support method %s", r.Method), 405)
		return
	}
	targetID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request body", 400)
		return
	}
	var updatedTodo models.Todo
	err = json.Unmarshal(body, &updatedTodo)
	if err != nil {
		http.Error(w, "Bad request body", 400)
		return
	}
	statusCode, err := x.repository.UpdateTodoByID(targetID, updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
	}
	w.WriteHeader(statusCode)
}
