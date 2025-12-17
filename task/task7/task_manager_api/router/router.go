package router

import (
	"taskManager/controllers"
	"taskManager/middleware"

	"github.com/gin-gonic/gin"
)

func Route()*gin.Engine{
	router:= gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"message": "pong"})
	})

	router.GET("/tasks", auth.AuthMiddleware(),controller.GetAllTasks)
	router.GET("/tasks/:id",auth.AuthMiddleware(), controller.GetTaskbyId)
	router.PUT("/tasks/:id",auth.AuthAdminMiddleware(), controller.UpdateTaskId)
	router.DELETE("/tasks/:id",auth.AuthAdminMiddleware(),controller.DeleteTaskById)
	router.POST("/tasks/new",auth.AuthAdminMiddleware(), controller.AddTask)
	
	router.POST("/register",controller.Register)
	router.POST("/login",controller.Login)
	router.POST("/promote/:username",auth.AuthAdminMiddleware(), controller.Promote)


	return router
}



//
