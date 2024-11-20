package models

import "time"

type Blog struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddBlog struct{
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
}