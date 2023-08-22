package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"

	"todoApi-go/pkg/domain/task"
)

type InMemoryTaskRepository struct {
	tasks map[uuid.UUID]task.Task //Todo: use Storage Task
	sync.Mutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[uuid.UUID]task.Task),
	}
}

func (tr *InMemoryTaskRepository) AddTask(t task.Task) error {

	if _, ok := tr.tasks[t.Id]; ok {
		return fmt.Errorf("Task already exists")
	}
	tr.Lock()
	tr.tasks[t.Id] = t
	tr.Unlock()
	return nil
}

func (tr *InMemoryTaskRepository) GetAllTasks() ([]task.Task, error) {
	if tr.tasks == nil {
		return nil, errors.New("tasks map not initialized")
	}
	return maps.Values(tr.tasks), nil
}

func (tr *InMemoryTaskRepository) GetTask(u uuid.UUID) (task.Task, error) {
	if task, ok := tr.tasks[u]; ok {
		return task, nil
	}

	return task.Task{}, errors.New("task not found")
}
