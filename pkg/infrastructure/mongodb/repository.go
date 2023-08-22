package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"todoApi-go/pkg/domain/task"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTaskRepository struct {
	db   *mongo.Database
	task *mongo.Collection
}

func NewMongoTaskRepository(ctx context.Context, connectionString string) (*MongoTaskRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("todo-db")
	tasks := db.Collection("task")

	return &MongoTaskRepository{
		db:   db,
		task: tasks,
	}, nil
}

func (tr *MongoTaskRepository) AddTask(task task.Task) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newTask := Create(task)
	_, err := tr.task.InsertOne(ctx, newTask)
	if err != nil {
		return err
	}
	return nil
}

func (tr *MongoTaskRepository) GetAllTasks() ([]task.Task, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cursor, err := tr.task.Find(ctx, bson.D{})

	var resultsTasks []MongoTask
	if err = cursor.All(ctx, &resultsTasks); err != nil {
		return []task.Task{}, err
	}

	var retrievedTasks []task.Task
	for _, result := range resultsTasks {
		t := result.ToDomain()
		retrievedTasks = append(retrievedTasks, t)
	}

	return retrievedTasks, nil
}

func (tr *MongoTaskRepository) GetTask(id uuid.UUID) (task.Task, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultTask := tr.task.FindOne(ctx, bson.M{"id": id})

	var retrievedTask MongoTask
	err := resultTask.Decode(&retrievedTask)
	if err != nil {
		return task.Task{}, err
	}
	return retrievedTask.ToDomain(), nil
}

func (m MongoTask) ToDomain() task.Task {
	// Create a ProxyTask
	t := task.Task{Id: m.Id, Title: m.Title, Completed: m.Completed}
	return t
}

func Create(t task.Task) *MongoTask {
	m := MongoTask{t.Id, t.Title, t.Completed, time.Now()}
	return &m
}
