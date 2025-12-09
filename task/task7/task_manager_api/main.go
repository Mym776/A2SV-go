package main

import (
	
	"taskManager/router"
	

)


// returns the entire list of tasks


func main() {

	router := router.Route()

	
	router.Run("localhost:3000")
}
