package services

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {
	
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("ERROR: HASHING_PWD")
	}
	return string(hashedPwd)
}

func ComparePassword(usrPwd string, hashedpwd string) bool{
	fmt.Println(usrPwd)
	fmt.Println(hashedpwd)
	err :=bcrypt.CompareHashAndPassword([]byte(hashedpwd), []byte(usrPwd))  
	if err != nil{
		return false
	}
	return true
}