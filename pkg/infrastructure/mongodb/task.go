package mongodb

import (
	"time"

	"github.com/google/uuid"
)

type MongoTask struct {
	Id        uuid.UUID `bson:"id"`
	Title     string    `bson:"name"`
	Completed bool      `bson:"completed"`
	Created   time.Time `bson:"created"`
}
