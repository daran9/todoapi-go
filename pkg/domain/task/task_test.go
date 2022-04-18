package task_test

import (
	"errors"
	"testing"
	"todoApi-go/pkg/domain/task"

	"github.com/stretchr/testify/assert"
)

func TestTask_NewTask(t *testing.T) {

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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Task
			_, err := task.NewTask(tc.arg, false)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
