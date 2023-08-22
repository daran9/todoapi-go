package task

import (
	"errors"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
}

func (t *Task) GetId() uuid.UUID {
	return t.Id
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) GetCompleted() bool {
	return t.Completed
}

func NewTask(title string, completed bool) (Task, error) {

	if title == "" {
		return Task{}, errors.New("empty title")
	}

	return Task{
		Id:        uuid.New(),
		Title:     title,
		Completed: completed,
	}, nil
}
