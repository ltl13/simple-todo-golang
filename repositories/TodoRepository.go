package repositories

import (
	"database/sql"
	"fmt"
	database "github.com/ltl13/simple-todo-golang/database"
	"github.com/ltl13/simple-todo-golang/models"
)

type TodoRepository struct {
}

func (t *TodoRepository) GetAllTodos() ([]models.Todo, int, error) {
	db := database.GetDBInstance()
	var todos []models.Todo
	rows, err := db.Query("SELECT * FROM TodoItem")
	if err != nil {
		return nil, 500, fmt.Errorf("getAllTodos: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Detail, &todo.IsDone); err != nil {
			return nil, 500, fmt.Errorf("getAllTodos: %v", err)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, 500, fmt.Errorf("getAllTodos: %v", err)
	}
	return todos, 200, nil
}

func (t *TodoRepository) GetTodoByID(id int64) (*models.Todo, int, error) {
	db := database.GetDBInstance()
	var todo *models.Todo = new(models.Todo)
	row := db.QueryRow("SELECT * FROM TodoItem WHERE id = ?", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Detail, &todo.IsDone); err != nil {
		if err == sql.ErrNoRows {
			return nil, 404, fmt.Errorf("getTodoByID %d: no such todo", id)
		}
		return nil, 500, fmt.Errorf("getTodoByID %d: %v", id, err)
	}
	return todo, 200, nil
}

func (t *TodoRepository) AddTodo(newTodo models.Todo) (int, error) {
	db := database.GetDBInstance()
	_, err := db.Exec(
		"INSERT INTO TodoItem (title, detail, isDone) VALUES (?, ?, ?)",
		newTodo.Title,
		newTodo.Detail,
		false,
	)
	if err != nil {
		return 500, fmt.Errorf("addTodo: %v", err)
	}
	return 201, nil
}

func (t *TodoRepository) RemoveTodoByID(id int64) (int, error) {
	db := database.GetDBInstance()
	todo, _, _ := t.GetTodoByID(id)
	if todo == nil {
		return 404, fmt.Errorf("removeTodoByID %d: no such todo", id)
	}
	_, err := db.Exec(
		`
		DELETE FROM TodoItem
		WHERE id = ?
		`,
		id,
	)
	if err != nil {
		return 500, fmt.Errorf("removeTodoByID %d: %v", id, err)
	}
	return 204, nil
}

func (t *TodoRepository) UpdateTodoByID(id int64, updatedTodo models.Todo) (int, error) {
	db := database.GetDBInstance()
	todo, _, _ := t.GetTodoByID(id)
	if todo == nil {
		return 404, fmt.Errorf("updatedTodoByID %d: no such todo", id)
	}
	_, err := db.Exec(
		`
		UPDATE TodoItem
		SET title = ?, detail = ?, isDone = ?
		WHERE id = ?
		`,
		updatedTodo.Title,
		updatedTodo.Detail,
		updatedTodo.IsDone,
		id,
	)
	if err != nil {
		return 500, fmt.Errorf("updateTodoByID %d: %v", id, err)
	}
	return 201, nil
}
