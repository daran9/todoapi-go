package task

import (
	"log"
	"net/http"

	"todoApi-go/pkg/domain/task"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	ts task.TaskService
}

type URI struct {
	Id string `uri:"id"`
}

func NewTaskController(s task.TaskService) *TaskController {
	return &TaskController{s}
}

// @Summary get all items in the task list
// @ID get-all-tasks
// @Produce json
// @Success 200 {object} task.Task
// @Router /task [get]
func (tc TaskController) GetTasks(c *gin.Context) {

	tasks, err := tc.ts.GetAllTasks()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, tasks)
}

// @Summary get a task item by ID
// @ID get-task-by-id
// @Produce json
// @Param id path string true "task ID"
// @Success 200 {object} task.Task
// @Failure 404 {object} object
// @Router /task/{id} [get]
func (tc TaskController) GetTaskById(c *gin.Context) {
	uri := URI{}

	// binding to URI
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := uuid.Parse(uri.Id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resTask, err := tc.ts.GetTask(id)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	c.JSON(http.StatusOK, resTask)
}

// @Summary add a new item to the task list
// @ID create-task
// @Produce json
// @Param data body task.Task true "task data"
// @Success 200 {object} task.Task
// @Failure 400 {object} object
// @Router /task [post]
func (tc TaskController) AddTask(c *gin.Context) {
	body := task.Task{}
	if err := c.BindJSON(&body); err != nil {
		log.Fatal(err)
	}
	id, err := tc.ts.AddTask(body)
	if err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
	}
	body.Id = id
	c.JSON(http.StatusAccepted, &body)
}

func (tc TaskController) DeleteTask(c *gin.Context) {
	//Todo: To be implemented
}
