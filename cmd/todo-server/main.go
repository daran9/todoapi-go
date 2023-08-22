package main

import (
	"context"
	"todoApi-go/pkg/controllers"
	"todoApi-go/pkg/controllers/task"
	task_domain "todoApi-go/pkg/domain/task"
	"todoApi-go/pkg/infrastructure/mongodb"
)

// @title 		Todo API
// @version 	1.1
// @description Todo API built using Go and Gin with DDD. You can visit the GitHub repository at https://github.com/daran9/todo-api-go

// @contact.name 	Sundaran Kumar
// @contact.url 	http://daran.wordpress.com
// @contact.email 	@daran9

// @license.name	MIT
// @license.url 	https://opensource.org/licenses/MIT

// @host 		localhost:8080
// @BasePath 	/api/v1
// @query.collection.format multi
func main() {

	taskRepository, err := mongodb.NewMongoTaskRepository(context.Background(), "mongodb://mongoadmin:secret123@mongo:27017")
	if err != nil {
		panic(err)
	}
	taskService := task_domain.NewAddTaskService(taskRepository)
	taskController := task.NewTaskController(taskService)

	r := controllers.Router(taskController)

	r.Run(":8081")
}
