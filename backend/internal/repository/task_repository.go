package repository

import "backend/internal/model"

type TaskRepository interface {
	GetAll() ([]model.Task, error)
	GetByID(id string) (model.Task, error)
	Create(task model.Task) (model.Task, error)
	Update(id string, task model.Task) (model.Task, error)
	Delete(id string) error
}