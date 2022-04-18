package controllers

import (
	"todoApi-go/pkg/controllers/task"

	"github.com/gin-gonic/gin"

	_ "todoApi-go/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Router(ctrl *task.TaskController) *gin.Engine {

	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("task", ctrl.GetTasks)
		v1.GET("task/:id", ctrl.GetTaskById)
		v1.POST("task", ctrl.AddTask)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
