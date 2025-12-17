package router

import (
	"taskManager/Delivery/controllers"
	services "taskManager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	// r.POST("/ping",func(ctx *gin.Context) {
	// 	data , _ := ctx.GetRawData()
		
	// 	ctx.IndentedJSON(200,string(data))
	// })
	// r.PUT("/register",controller.Register)
	r.GET("/AllTasks",services.AuthMiddleware(),controller.GetAllTasks)
	r.GET("/Tasks/:id",services.AuthMiddleware(),controller.GetTaskByID)
	r.POST("/AddTask", services.AuthAdminMiddleware(),controller.AddTask)
	r.PUT("/UpdateTask/:id", services.AuthAdminMiddleware(),controller.UpdateTask)
	r.DELETE("/DeleteTask/:id",services.AuthAdminMiddleware(),controller.DeleteTask)

	r.POST("/Register", controller.Register)
	r.POST("/Login", controller.Login)
	r.POST("/Promote/:username",services.AuthAdminMiddleware(),controller.Promote)
}