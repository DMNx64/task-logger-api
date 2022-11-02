package main

import "github.com/google/uuid"

type task struct {
	ID uuid.UUID `json:"id"`
	Title string `json:"title"`
}