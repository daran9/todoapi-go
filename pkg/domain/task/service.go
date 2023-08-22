package task

import "github.com/google/uuid"

type TaskService interface {
	AddTask(Task) (uuid.UUID, error)
	GetAllTasks() ([]Task, error)
	GetTask(uuid.UUID) (Task, error)
}

type taskService struct {
	r TaskRepository
}

func NewAddTaskService(r TaskRepository) TaskService {
	return &taskService{r}
}

func (ts *taskService) AddTask(reqTask Task) (uuid.UUID, error) {

	newTask, err := NewTask(reqTask.Title, reqTask.Completed)
	if err != nil {
		return uuid.UUID{}, err
	}
	if errAdd := ts.r.AddTask(newTask); errAdd != nil {
		return uuid.UUID{}, errAdd
	}

	return newTask.Id, nil
}

func (ts *taskService) GetAllTasks() ([]Task, error) {
	return ts.r.GetAllTasks()
}

func (ts *taskService) GetTask(u uuid.UUID) (Task, error) {
	return ts.r.GetTask(u)
}
