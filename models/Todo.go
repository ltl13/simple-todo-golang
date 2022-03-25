package models

type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	IsDone bool   `json:"isDone"`
}
