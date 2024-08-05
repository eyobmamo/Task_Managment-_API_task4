package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	R := gin.Default()

	RouteControl := *controllers.NewTaskController()

	R.GET("/tasks", RouteControl.GetAllTasks)
	R.POST("/tasks", RouteControl.CreateTasks)
	R.GET("/tasks/:id", RouteControl.GetTasksById)
	R.PUT("/tasks/:id", RouteControl.UpdateTasksById)
	R.DELETE("/tasks/:id", RouteControl.DeleteTasksById)

	return R

}
