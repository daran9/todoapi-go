package memory

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id        uuid.UUID
	Title     string
	Completed bool
	Created   time.Time
}
