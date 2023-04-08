package task

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Created   time.Time `json:"Created"`
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

func (t *Task) SetCreated(created time.Time) {
	t.Created = created
}

func (t *Task) SetID(id uuid.UUID) {
	t.Id = id
}

func NewTask(title string, completed bool) (Task, error) {

	if title == "" {
		return Task{}, errors.New("empty title")
	}

	return Task{
		Id:        uuid.New(),
		Title:     title,
		Completed: completed,
		Created:   time.Now(),
	}, nil
}
