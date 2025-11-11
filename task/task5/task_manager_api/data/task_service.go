package service

import (
	
	"time"

	"taskManager/models"
)


// Task data
var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Completed"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "In Progress"},
	{ID: "4", Title: "Task 4", Description: "Fourth task", DueDate: time.Now().AddDate(0,0,3), Status: "In Progress"},
	{ID: "5", Title: "Task 5", Description: "Fifth task", DueDate: time.Now().AddDate(0, 0, 4), Status: "In Progress"},
	{ID: "6", Title: "Task 6", Description: "Sixth task", DueDate: time.Now().AddDate(0, 0, 5), Status: "In Progress"},
}


func Tasks() []models.Task{
	return tasks
}


func TaskId(id string) models.Task {
	j := models.Task{}
	for _, i := range tasks {
		if i.ID == id {
			j=i
			break
		}
	}
	return j
}

func UpdateTask(id string, update models.Task) bool {
	
	for i, task := range tasks {
		if task.ID == id {
				tasks[i] = update
				return true
		}
	}
	return false
	
}

func DeleteTask(id string) bool {

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			return true
		}
	}
	return false
}


func AddTask(task models.Task) bool {
	for _, t := range tasks {
		if t.ID == task.ID {
			return false
		}
	}

	tasks = append(tasks, task)
	return true
}