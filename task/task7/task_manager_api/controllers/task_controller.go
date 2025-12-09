package controller

import (
	// "fmt"
	// "log"
	"net/http"
	// "time"

	"taskManager/data"
	"taskManager/models"

	"github.com/gin-gonic/gin"

	
)

func GetAllTasks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, service.Tasks(c))
}

// returns a task with a prespecified id
func GetTaskbyId(c *gin.Context) {
	id := c.Param("id")
	task := service.TaskId(c, id)
	if (task == models.Task{}) {
		c.IndentedJSON(404, gin.H{"message":"file not found"})

	} else {
		c.IndentedJSON(http.StatusOK, task)
	}
}

// updates a task with a prespecified id
func UpdateTaskId(c *gin.Context) {
	id := c.Param("id")
	var update models.Task

	if err := c.ShouldBindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid task"})
		return
	}

	stat := service.UpdateTask(c,id, update)

	if stat {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated successfully"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})

}

// deletes  a task with a prespecified id

func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")

	stat := service.DeleteTask(c,id)
	if stat {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})

}

func AddTask(c *gin.Context) {

	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid task"})
		return
	}

	stat := service.AddTask(c,task)

	if stat {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task added successfully"})
	}else{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "task already available"})
	}

}

func Register(c *gin.Context){
	service.Register(c)
}
func Login(c *gin.Context){
	service.Login(c)
}
func Promote (c *gin.Context){
	var username = c.Param("username")

	service.Promote(c,username)
}