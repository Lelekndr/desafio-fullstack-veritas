package model

import (
	"time"
)

type Task struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	DateCreation time.Time `json:"date_creation"`
	Importance   string    `json:"importance"`
	Column       string    `json:"column"`
}