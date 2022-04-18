package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"

	"todoApi-go/pkg/domain/task"
)

type TaskRepository struct {
	tasks map[uuid.UUID]task.Task //Todo: use Storage Task
	sync.Mutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[uuid.UUID]task.Task),
	}
}

func (tr *TaskRepository) AddTask(t task.Task) error {

	if _, ok := tr.tasks[t.Id]; ok {
		return fmt.Errorf("Task already exists")
	}
	tr.Lock()
	tr.tasks[t.Id] = t
	tr.Unlock()
	return nil
}

func (tr *TaskRepository) GetAllTasks() ([]task.Task, error) {
	if tr.tasks == nil {
		return nil, errors.New("tasks map not initialized")
	}
	return maps.Values(tr.tasks), nil
}

func (tr *TaskRepository) GetTask(u uuid.UUID) (task.Task, error) {
	if task, ok := tr.tasks[u]; ok {
		return task, nil
	}

	return task.Task{}, errors.New("task not found")
}
