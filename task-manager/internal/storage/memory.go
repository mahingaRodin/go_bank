package storage

import (
	"errors"
	"sync"
	"task-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)


type TaskStorage interface {
	Create(task *models.Task) error
	GetByID(id string) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	Update(id string, update *models.TaskUpdate) (*models.Task, error)
	Delete(id string) error
	GetByStatus(status string) ([]*models.Task, error)
}

type MemoryStorage struct {
	tasks map[string]*models.Task
	mutex sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[string]*models.Task),
	}
}


//adds a new task to the in-memory storage
func(s *MemoryStorage) Create(task *models.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	s.tasks[task.ID] = task
	return nil
}


//getById
func(s *MemoryStorage) GetByID(id string) (*models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()


	task, exists := s.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}

	return task, nil
}

//retrieves all tasks 
func(s *MemoryStorage) GetAll() ([]*models.Task, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	tasks := make([]*models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

//udpating a task 
func(s *MemoryStorage) Update(id string, update *models.TaskUpdate) (*models.Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}
	
	//uypdate the provided fields
	if update.Title != nil {
		task.Title = *update.Title
	}

	if update.Description != nil {
		task.Description = *update.Description
	}

	if update.Status != nil {
		task.Status = *update.Status
	}

	if update.Priority != nil {
		task.Priority = *update.Priority
	}

	if update.DueDate != nil {
		task.DueDate = *update.DueDate
	}

	task.UpdatedAt = time.Now()

	return task, nil
}

// Delete removes a task from storage
func (s *MemoryStorage) Delete(id string) error {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    if _, exists := s.tasks[id]; !exists {
        return ErrTaskNotFound
    }

    delete(s.tasks, id)
    return nil
}

// GetByStatus retrieves tasks by their status
func (s *MemoryStorage) GetByStatus(status string) ([]*models.Task, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()

    var filteredTasks []*models.Task
    for _, task := range s.tasks {
        if task.Status == status {
            filteredTasks = append(filteredTasks, task)
        }
    }

    return filteredTasks, nil
}


