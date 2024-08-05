package controllers

import (
	"fmt"
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"task_manager/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type RouteController interface {
	GetAllTasks(c *gin.Context)
	CreateTasks(c *gin.Context)
	GetTasksById(c *gin.Context)
	UpdateTasksById(c *gin.Context)
	DeleteTasksById(c *gin.Context)
}

type RouteControl struct {
	taskService data.TaskService
}

func NewTaskController() *RouteControl {
	return &RouteControl{
		taskService: *data.NewTaskService(),
	}
}

func (tc *RouteControl) GetAllTasks(c *gin.Context) {

	tasks, err := tc.taskService.GetTasks()
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)

}

func (tc *RouteControl) CreateTasks(c *gin.Context) {
	var newTak models.Task
	v := validator.New()
	if err := c.ShouldBindJSON(&newTak); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid data", "error": err.Error()})
		return
	}
	if err := v.Struct(newTak); err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}
	err := utils.ValidateStatus(&newTak)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
	}

	//
	task, err := tc.taskService.CreateTasks(&newTak)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, task)
	}
}

func (tc *RouteControl) GetTasksById(c *gin.Context) {
	id := c.Param("id")

	task, err := tc.taskService.GetTasksById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(200, task)
	}
}

func (tc *RouteControl) UpdateTasksById(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//
	data, er := tc.taskService.UpdateTasksById(id, updatedTask)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Task updated", "data": data})
	}

}

func (tc *RouteControl) DeleteTasksById(c *gin.Context) {
	id := c.Param("id")

	err := tc.taskService.DeleteTasksById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Data Deleted"})
}
