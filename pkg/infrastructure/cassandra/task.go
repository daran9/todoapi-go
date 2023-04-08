package cassandra

import (
	"time"
	"todoApi-go/pkg/domain/task"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type TaskItem struct {
	Id        gocql.UUID
	Title     string
	Completed bool
	Created   time.Time
}

func NewFromTask(t task.Task) TaskItem {
	return TaskItem{
		Id:        gocql.UUID(t.Id),
		Title:     t.Title,
		Completed: t.Completed,
		Created:   t.Created,
	}
}

func (ti TaskItem) ToTask() task.Task {
	t, _ := task.NewTask(ti.Title, ti.Completed)
	t.SetID(uuid.UUID(ti.Id))
	t.SetCreated(ti.Created)
	return t
}
