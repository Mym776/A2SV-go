package main

import (
	// "fmt"
	// "library/models"
	"library/controllers"
	"library/services"
)

func main() {
	// var library = services.Library{make(map[int]models.Book),make(map[int]models.Member)};
	// book:= models.Book{1,"Harry potter","JK Rowling","available"}
	// member:= models.Member{ID: 1,Name:"Parry Hoter", BorrowedBooks: []models.Book{}}
	// library.AddBook(book)
	// library.MemberList[member.ID]=member
	// library.BorrowBook(book.ID,member.ID)
	// fmt.Println(library)
	// t := library.ListAvailableBooks()
	l := services.NewLibrary()

	controllers.Start(l)

}
