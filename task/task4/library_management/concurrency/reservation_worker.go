package concurrency

import (
	"library/services"
	
)

func Reservation(memberID int , bookID int, l services.LibraryManager){

	go func (){
		l.ReserveBook(memberID,bookID)
	}()
	
	
	go func(){
		l.UnReserveBook(memberID,bookID)
	}()
	return
	
	

}