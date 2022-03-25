package main

import (
	controller "github.com/ltl13/simple-todo-golang/controllers"
	"log"
	"net/http"
)

func main() {
	var todoController *controller.TodoController = new(controller.TodoController)
	http.HandleFunc("/api/todo/get-all", todoController.GetAllTodos)
	http.HandleFunc("/api/todo/get", todoController.GetTodoByID)
	http.HandleFunc("/api/todo/add", todoController.AddTodo)
	http.HandleFunc("/api/todo/update", todoController.UpdateTodoByID)
	http.HandleFunc("/api/todo/remove", todoController.RemoveTodoByID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
