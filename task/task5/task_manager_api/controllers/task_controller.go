package controller

import (
	"net/http"

	"taskManager/data"
	"taskManager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, service.Tasks())
}

// returns a task with a prespecified id
func GetTaskbyId(c *gin.Context) {
	id := c.Param("id")
	task := service.TaskId(id)
	if (task == models.Task{}) {
		c.IndentedJSON(http.StatusNotFound, nil)

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

	stat := service.UpdateTask(id, update)

	if stat {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task updated successfully"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})

}

// deletes  a task with a prespecified id

func DeleteTaskById(c *gin.Context) {
	id := c.Param("id")

	stat := service.DeleteTask(id)
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

	stat := service.AddTask(task)

	if stat {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "task added successfully"})
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "task already available"})

}
