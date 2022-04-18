package task_test

import (
	"errors"
	"testing"
	"todoApi-go/pkg/domain/task"
	"todoApi-go/pkg/domain/task/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTask_AddTask(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		name        string
		arg         string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			name:        "Empty Name validation",
			arg:         "",
			expectedErr: errors.New("empty title"),
		}, {
			name:        "Valid Name",
			arg:         "Nick Chavas",
			expectedErr: nil,
		},
	}

	repo := &mocks.AddTaskRepository{}
	repo.On("AddTask", mock.AnythingOfType("task.Task")).
		Return(nil).
		Once()

	service := task.NewAddTaskService(repo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Create a new Task
			task := task.Task{
				Title: tc.arg,
			}
			_, err := service.AddTask(task)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestTask_GetTask(t *testing.T) {

	addedTask := task.Task{
		Title: "title",
	}

	repo := &mocks.AddTaskRepository{}
	repo.On("GetTask", mock.AnythingOfType("uuid.UUID")).
		Return(addedTask, nil).
		Once()

	service := task.NewAddTaskService(repo)

	_, err := service.GetTask(uuid.New())

	assert.Equal(t, nil, err)

}

func TestTask_GetAllTasks(t *testing.T) {

	addedTask := task.Task{
		Title: "title",
	}

	repo := &mocks.AddTaskRepository{}
	repo.On("GetAllTasks").
		Return([]task.Task{addedTask}, nil).
		Once()

	service := task.NewAddTaskService(repo)

	_, err := service.GetAllTasks()

	assert.Equal(t, nil, err)
}
