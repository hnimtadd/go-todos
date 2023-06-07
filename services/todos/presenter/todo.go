package presenter

import "time"

type TodoResponse struct {
	Id        string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoRequest struct {
	Content string
}
