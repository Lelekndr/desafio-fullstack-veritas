package repository

import (
	"encoding/json"
	"os"
	"sync"
	"time"
	"github.com/google/uuid"
	"backend/internal/model"
	"fmt"
)

type JSONTaskRepository struct {
	filePath string
	mu       sync.Mutex
	tasks    map[string]model.Task
}

func NewJSONTaskRepository(filePath string) (*JSONTaskRepository, error) {
	repo := &JSONTaskRepository{
		filePath: filePath,
		tasks:    make(map[string]model.Task),
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return repo, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, &repo.tasks); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *JSONTaskRepository) saveToFile() error {
	data, err := json.MarshalIndent(r.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

func (r *JSONTaskRepository) Create(task model.Task) (model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	task.ID = uuid.New().String()
	task.DateCreation = time.Now()
	task.Column = "to_do"

	r.tasks[task.ID] = task
	if err := r.saveToFile(); err != nil {
		return model.Task{}, err
	}

	return task, nil
}

func (r *JSONTaskRepository) GetAll() ([]model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := make([]model.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *JSONTaskRepository) GetByID(id string) (model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return model.Task{}, fmt.Errorf("task with ID %s not found", id)
	}
	return task, nil
}

func (r *JSONTaskRepository) Update(id string, updatedTask model.Task) (model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existingTask, ok := r.tasks[id]
	if !ok {
		return model.Task{}, fmt.Errorf("task with ID %s not found", id)
	}

	updatedTask.ID = existingTask.ID
	updatedTask.DateCreation = existingTask.DateCreation

	r.tasks[id] = updatedTask
	if err := r.saveToFile(); err != nil {
		return model.Task{}, err
	}

	return updatedTask, nil
}

func (r *JSONTaskRepository) Delete(id string) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[id]
	if !ok {
		return fmt.Errorf("task with ID %s not found", id)
	}
	delete(r.tasks, id)
	return r.saveToFile()
}