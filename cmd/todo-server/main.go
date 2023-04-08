package main

import (
	"todoApi-go/pkg/controllers"
	"todoApi-go/pkg/controllers/task"
	task_domain "todoApi-go/pkg/domain/task"
	"todoApi-go/pkg/infrastructure/cassandra"
)

// @title 		Todo API
// @version 	1.0
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

	//taskRepository := memory.NewTaskRepository()
	taskRepository := cassandra.NewTaskRepository("cass1")
	taskService := task_domain.NewAddTaskService(taskRepository)
	taskController := task.NewTaskController(taskService)

	r := controllers.Router(taskController)

	r.Run(":8081")
}
