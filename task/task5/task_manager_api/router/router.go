package router

import (
	 "taskManager/controllers"

	"github.com/gin-gonic/gin"
)

func Route()*gin.Engine{
	router:= gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"message": "pong"})
	})

	router.GET("/tasks", controller.GetAllTasks)
	router.GET("/tasks/:id", controller.GetTaskbyId)
	router.PUT("/tasks/:id", controller.UpdateTaskId)
	router.DELETE("/tasks/:id",controller.DeleteTaskById)
	router.POST("/tasks/new", controller.AddTask)
	return router
}

