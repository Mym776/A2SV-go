package main

import (
	// "fmt"
	// "library/models"
	"library/controllers"
	"library/services"
)

func main() {
	
	l := services.NewLibrary()

	controllers.Start(l)

}
