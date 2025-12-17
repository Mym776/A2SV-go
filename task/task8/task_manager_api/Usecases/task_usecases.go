package usecases

import (
	entities "taskManager/Domain"
	repositories "taskManager/Repositories"

	
)

type TaskUsecase struct {
	Repo repositories.TaskRepository
}

func(t *TaskUsecase) GetAllTasks() []entities.Task{
	
	var tasks = t.Repo.GetAllTask()
	
	return tasks
}

func(t *TaskUsecase) GetTaskByID(id string) entities.Task{

	var task = t.Repo.GetTaskByID(id)
	return task 
}

func(t *TaskUsecase) AddTask(task entities.Task) bool{
	
	success := t.Repo.AddTask(task)
	return success
}


func(t *TaskUsecase) UpdateTask(task entities.Task, id string) bool{
	success := t.Repo.UpdateTask(task, id)
	return success
}

func(t *TaskUsecase) DeleteTask(id string) bool{
	success := t.Repo.DeleteTask(id)
	return success
}