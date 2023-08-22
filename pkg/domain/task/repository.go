package task

import "github.com/google/uuid"

type TaskRepository interface {
	AddTask(Task) error
	GetAllTasks() ([]Task, error)
	GetTask(uuid.UUID) (Task, error)
}
