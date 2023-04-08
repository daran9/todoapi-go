package cassandra

import (
	"fmt"
	"time"
	"todoApi-go/pkg/domain/task"

	"github.com/avast/retry-go"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type TaskRepository struct {
	Session  *gocql.Session
	hostname string
}

func NewTaskRepository(hostname string) *TaskRepository {
	tr := TaskRepository{}
	tr.hostname = hostname

	err := retry.Do(
		func() error {
			err := tr.init()
			if err != nil {
				return err
			}
			return nil
		},
		retry.Attempts(20),
		retry.Delay(200),
	)
	if err != nil {
		panic(err)
	}
	return &tr
}

func (tr *TaskRepository) init() error {
	var err error
	cluster := gocql.NewCluster(tr.hostname)
	cluster.ProtoVersion = 4
	cluster.Keyspace = "tododb"
	tr.Session, err = cluster.CreateSession()
	if err != nil {
		return err
	}
	fmt.Println("cassandra well initialized")
	return nil
}

func (tr *TaskRepository) AddTask(t task.Task) error {

	if err := tr.Session.Query("INSERT INTO tododb.tasks_by_id(id, title, completed, created) VALUES(?, ?, ?, ?)",
		t.Id, t.Title, t.Completed, t.Created).Exec(); err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) GetAllTasks() ([]task.Task, error) {

	var taskItems []TaskItem
	m := map[string]interface{}{}

	iter := tr.Session.Query("SELECT * FROM tododb.tasks_by_id").Iter()
	for iter.MapScan(m) {
		taskItems = append(taskItems, TaskItem{
			Id:        m["id"].(gocql.UUID),
			Title:     m["title"].(string),
			Completed: m["completed"].(bool),
			Created:   m["created"].(time.Time),
		})
		m = map[string]interface{}{}
	}
	tasks := make([]task.Task, len(taskItems))
	for i, v := range taskItems {
		tasks[i] = v.ToTask()
	}
	return tasks, nil
}

func (tr *TaskRepository) GetTask(taskId uuid.UUID) (task.Task, error) {

	var tasks []TaskItem
	m := map[string]interface{}{}

	iter := tr.Session.Query("SELECT * FROM tododb.tasks_by_id WHERE id=?", taskId).Iter()
	for iter.MapScan(m) {
		tasks = append(tasks, TaskItem{
			Id:        m["id"].(gocql.UUID),
			Title:     m["title"].(string),
			Completed: m["completed"].(bool),
			Created:   m["created"].(time.Time),
		})
		m = map[string]interface{}{}
	}
	return tasks[0].ToTask(), nil
}
