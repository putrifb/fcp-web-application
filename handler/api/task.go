package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	task := c.Param("id")
	id, err := strconv.Atoi(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("Invalid task ID")) //400
		return
	}
	var upTask = model.Task{}

	if err = c.BindJSON(&upTask); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("Something went wrong while updating task!")) //500
		return
	}
	upTask.ID = id
	if err = t.taskService.Update(upTask.ID, &upTask); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("Something went wrong while updating task!")) //500
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResponse("update task success")) //200// TODO: answer here
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	task := c.Param("id")
	id, err := strconv.Atoi(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("Invalid task ID")) //400
		return
	}
	if err = t.taskService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("Something went wrong while updating task!")) //500
		return
	}
	c.JSON(http.StatusOK, model.NewSuccessResponse("delete task success")) //200
	// TODO: answer here
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	//mendapatkan daftar tugas.
	list, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()}) //500
		return
	}
	c.JSON(http.StatusOK, list) //200
	// TODO: answer here
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	//mendapatkan daftar tugas dengan nama kategorinya.
	task := c.Param("id")
	id, err := strconv.Atoi(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("Invalid task ID")) //400
		return
	}
	list, err := t.taskService.GetTaskCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()}) //500
		return
	}
	c.JSON(http.StatusOK, list) //200
	// TODO: answer here
}
