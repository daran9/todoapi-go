package memory

import (
	"errors"
	"testing"
	"todoApi-go/pkg/domain/task"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTaskRepository_AddTask(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	// Create a new Task
	task := task.Task{
		Title: "addTask",
		Id:    uuid.New(),
	}

	err := repo.AddTask(task)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(repo.tasks), "Expected 1 task added")
}

func TestTaskRepository_GetTask(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	// Create a new Task
	expectedTask := task.Task{
		Title: "getTask",
		Id:    uuid.New(),
	}

	err := repo.AddTask(expectedTask)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(repo.tasks), "Expected 1 task added")

	// Build our needed testcase data struct
	type testCase struct {
		name        string
		arg         uuid.UUID
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			name:        "Get Task by id",
			arg:         expectedTask.Id,
			expectedErr: nil,
		}, {
			name:        "Get Task by non-existent id",
			arg:         uuid.New(),
			expectedErr: errors.New("task not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Get Task
			_, err := repo.GetTask(tc.arg)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestTaskRepository_GetAllTasks(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	// Create a new Task
	task := task.Task{
		Title: "getAllTasks",
		Id:    uuid.New(),
	}

	err := repo.AddTask(task)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, len(repo.tasks), "Expected 1 task added")

	tasks, err := repo.GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(tasks), "Expected 1 task")
}
