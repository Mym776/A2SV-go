package usecases

import (
	entities "taskManager/Domain"
	repositories "taskManager/Repositories"
)

type UserUsecase struct{
	Repo repositories.UserRepository
}


func(t *UserUsecase) Register(user entities.User) bool{
	var success = t.Repo.Register(user)
	return success
}
func(t *UserUsecase) Login(user entities.User) bool{
	var success = t.Repo.Login(user)
	return success
}

func(t *UserUsecase) Promote(username string) bool{
	var success = t.Repo.Promote(username)
	return success
}