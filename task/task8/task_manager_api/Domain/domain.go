package entities

import "time"

type User struct {
	Name	string	`json:"name"`
	PasswordHash	string	`json:"passwordHash"`
	Role	string	`json:"role"`
}

type Task struct {
	ID	string	`json:"id"`
	Title	string	`json:"title"`
	Description	string	`json:"description"`
	DueDate	time.Time	`json:"due_date"`
	Status	string	`json:"status"`
}