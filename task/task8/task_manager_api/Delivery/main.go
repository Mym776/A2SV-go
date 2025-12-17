package main

import (
	router "taskManager/Delivery/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	rtr := gin.Default()
	router.Routes(rtr)
	rtr.Run()
}