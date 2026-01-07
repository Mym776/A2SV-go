package controller

import (
	"fmt"
	entities "taskManager/Domain"
	repositories "taskManager/Repositories"
	usecases "taskManager/Usecases"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {

	t := usecases.TaskUsecase{Repo: repositories.TaskRepository{Context: c}}
	tasks := t.GetAllTasks()
	if len(tasks)==0{
		c.IndentedJSON(204,gin.H{"message":"No Tasks available"})
		return
	}
	c.IndentedJSON(200,gin.H{"Tasks":tasks})

}

func GetTaskByID(c *gin.Context)  {
	t := usecases.TaskUsecase{Repo: repositories.TaskRepository{Context: c}}
	id := c.Param("id")
	task := t.GetTaskByID(id)
	if task == (entities.Task{}){
		c.IndentedJSON(404,gin.H{"message":"No Task with specified ID"})
	}
	c.IndentedJSON(200,gin.H{"Task":task})

}

func AddTask(c *gin.Context){
	t := usecases.TaskUsecase{Repo: repositories.TaskRepository{Context: c}}
	var task entities.Task
	
	err := c.ShouldBind(&task)
	if err != nil{
		c.IndentedJSON(400, gin.H{"error": "Invalid format"})	
		return
	}
	success := t.AddTask(task)
	if success{
		c.IndentedJSON(200,gin.H{"message": "Task added Successfully"})
		return
	}else{
		c.IndentedJSON(406,gin.H{"message": "Task Already exists"})
	}
}


func UpdateTask(c *gin.Context){
	t := usecases.TaskUsecase{Repo: repositories.TaskRepository{Context: c}}
	var task entities.Task
	
	err := c.ShouldBind(&task)
	if err != nil{
		c.IndentedJSON(400, gin.H{"error": "Invalid format"})	
		return
	}
	id := c.Param("id")
	success := t.UpdateTask(task,id)
	if success{
		c.IndentedJSON(200,gin.H{"message": "Task updated successfully"})
		return
	}
	c.IndentedJSON(400,gin.H{"error": "Task update failed"})

}


func DeleteTask(c *gin.Context){
	t := usecases.TaskUsecase{Repo: repositories.TaskRepository{Context: c}}
	id := c.Param("id")
	success := t.DeleteTask((id))
	if success {
		c.IndentedJSON(200 , gin.H{"message": "task deleted successfully"})
		return
	}
	c.IndentedJSON(406 , gin.H{"error": "task not deleted"})
}



func Register(c *gin.Context){
	u :=usecases.UserUsecase{Repo: repositories.UserRepository{Context: c}}
	
	var user entities.User
	
	err := c.ShouldBind(&user)
	fmt.Println(user)
	if err != nil{
		c.IndentedJSON(400,gin.H{"error": "Invalid format"})
		return
	}

	success := u.Register(user)
	
	if success {
		c.IndentedJSON(200,gin.H{"message":"user registered successfully"})
		return
	}
	c.IndentedJSON(400,gin.H{"error":"User not registered"})
}

func Login(c *gin.Context){
	u :=usecases.UserUsecase{Repo: repositories.UserRepository{Context: c}}
	var user entities.User

	err := c.ShouldBind(&user)

	if err != nil{
		c.IndentedJSON(400,gin.H{"error": "Invalid format"})
		return
	}
	success := u.Login(user)
	if success {
		return
	}
	c.IndentedJSON(400, gin.H{"error":"login failed"})

}

func Promote(c *gin.Context){
	u :=usecases.UserUsecase{Repo: repositories.UserRepository{Context: c}}
	username := c.Param("username")

	success := u.Promote(username)

	if success {
		c.IndentedJSON(200,gin.H{"message":"user Promoted successfully"})
		return
	}
	c.IndentedJSON(400,gin.H{"message":"Promotion failed"})
}