package mongodb

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"testing"
	"todoApi-go/pkg/domain/task"
)

var mongoServer *memongo.Server

func TestMain(m *testing.M) {
	var err error
	mongoServer, err = memongo.StartWithOptions(&memongo.Options{MongoVersion: "5.0.20", DownloadURL: "https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-rhel80-5.0.20.tgz"})
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()

	os.Exit(m.Run())
}

func TestTaskRepository_AddTask(t *testing.T) {
	repo, err := NewMongoTaskRepository(context.TODO(), mongoServer.URI())
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Task
	task := task.Task{
		Title: "addTask",
		Id:    uuid.New(),
	}

	addTaskErr := repo.AddTask(task)
	if addTaskErr != nil {
		t.Error(addTaskErr)
	}

	var retrievedTask MongoTask
	resultTasks := repo.task.FindOne(context.TODO(), bson.M{"id": task.Id})
	taskErr := resultTasks.Decode(&retrievedTask)
	assert.Nil(t, taskErr)
	assert.EqualValues(t, task.Id, retrievedTask.Id)
	assert.EqualValues(t, task.Title, retrievedTask.Title)
}

func TestTaskRepository_GetTask(t *testing.T) {
	repo, err := NewMongoTaskRepository(context.TODO(), mongoServer.URI())
	if err != nil {
		t.Error(err)
	}

	// Create a new Task
	expectedTask := task.Task{
		Title: "getTask",
		Id:    uuid.New(),
	}
	newTask := Create(expectedTask)
	_, insertErr := repo.task.InsertOne(context.TODO(), newTask)
	if insertErr != nil {
		t.Error(insertErr)
	}

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
		},
		{
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
	repo, err := NewMongoTaskRepository(context.TODO(), mongoServer.URI())
	if err != nil {
		log.Fatal(err)
	}

	// Create new slice of Tasks
	expectedTasks := [2]task.Task{
		{
			Title: "getAllTasks1",
			Id:    uuid.New(),
		},
		{
			Title: "getAllTasks2",
			Id:    uuid.New(),
		},
	}
	mangoTasks := make([]interface{}, 2)
	for _, et := range expectedTasks {
		newTask := Create(et)
		mangoTasks = append(mangoTasks, newTask)
	}
	_, insertErr := repo.task.InsertMany(context.TODO(), mangoTasks)
	if insertErr != nil {
		t.Error(insertErr)
	}

	tasks, err := repo.GetAllTasks()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(tasks), "Expected 1 task")
}
