package entity

import ("time")

type Task struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	AssignName string `json:"assign_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AssignName  string `json:"assign_name"`
}

