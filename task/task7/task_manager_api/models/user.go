package models


type User struct {
//	ID          string    `json:"id"`
	Name       string    `json:"name"`
	PasswordHash	string		`json:"passwordHash"`
	Role	string    `json:"role"`
} 