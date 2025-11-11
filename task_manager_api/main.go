package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

// define task
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Task data
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Completed"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "In Progress"},
	{ID: "4", Title: "Task 4", Description: "Fourth task", DueDate: time.Now().AddDate(0,0,3), Status: "In Progress"},
	{ID: "5", Title: "Task 5", Description: "Fifth task", DueDate: time.Now().AddDate(0, 0, 4), Status: "In Progress"},
	{ID: "6", Title: "Task 6", Description: "Sixth task", DueDate: time.Now().AddDate(0, 0, 5), Status: "In Progress"},
}

// returns the entire list of tasks
func getAllTasks(c *gin.Context) {
	
	c.IndentedJSON(http.StatusOK, tasks)
}


// returns a task with a prespecified id
func getTaskbyId(c *gin.Context) {
	id := c.Param("id")
	for _, i := range tasks {
		if i.ID == id {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, nil)
}


// updates a task with a prespecified id
func updateTaskId(c *gin.Context) {
	id := c.Param("id")

	var update Task

	if err := c.ShouldBindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid task"})
		return
	}
	for i, task := range tasks {
		if task.ID == id {

			if update.Title != "" {
				tasks[i].Title = update.Title
			}

			if update.Description != "" {
				tasks[i].Description = update.Description
			}

			c.IndentedJSON(http.StatusOK, gin.H{"message":"task updated successfully"})
			return
		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})

}

// deletes  a task with a prespecified id

func deleteTaskById(c *gin.Context){

	id := c.Param("id")
	for i, task := range tasks{
		if task.ID==id{
			tasks = append(tasks[:i],tasks[i+1:]...)		
			
			c.IndentedJSON(http.StatusOK, gin.H{"message":"task deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})

}


func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"message": "pong"})
	})

	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTaskbyId)
	router.POST("/tasks/:id", updateTaskId)
	router.DELETE("/tasks/:id",deleteTaskById)
	router.Run("localhost:3000")
}
