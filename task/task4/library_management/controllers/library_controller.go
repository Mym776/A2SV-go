package controllers

import (
	"bufio"
	"fmt"
	"library/models"
	"library/services"
	"log"
	"os"
	"strconv"
	"strings"
	// "strconv"
	// "strings"
)

// library_contoller.go accepts user input and calls on appropriate functions from library_service.go

func Start(l services.LibraryManager) {

	menu()

	for {
		var inp int
	fmt.Println("\n - - - - - - - - - - - - - - - - ")
		
		fmt.Print("Please input from 1-7: ")
		
		inp = inputInt()
		switch inp {
		case 1:
			//Add a book
			addbook(l)

		case 2:
			//Add a member
			addmember(l)

		case 3:
			//Remove a book
			removebook(l)

		case 4:
			//Borrow a book
			borrowbook(l)

		case 5:
			//Return a book
			returnbook(l)

		case 6:
			//List available books
			listavailablebooks(l)

		case 7:
			//List borrowed books by a member
			listborrowedbook(l)

		case -1:
			// exit
			return

		default:
			fmt.Println("Invalid choice")

		}
	}

}

// displays menu
func menu() {
	fmt.Println("Welcome to the Library")
	fmt.Println("1. Add a book")
	fmt.Println("2. Add a Member")
	fmt.Println("3. Remove a book")
	fmt.Println("4. Borrow a book")
	fmt.Println("5. Return a book")
	fmt.Println("6. List all available books")
	fmt.Println("7. List all borrowed books by a member")
	fmt.Println("-1. Exit")

}


//accepts book name and book author, creates a book instance and adds it to the library
func addbook(l services.LibraryManager) {
	fmt.Println(" - - - - - - - - - - - - - - - - ")

	var title string
	var author string
	var status string = "available"

	for {

		fmt.Println("Book Title:")
		title = input()
		


		fmt.Println("Book author:")
		author = input()

		// title = strings.TrimSpace(title)
		// author = strings.TrimSpace(author)
		if title!="" && author !="" {
			
			break
		}
		fmt.Println("! MISSING AN ENTRY !")
	}
	var book = models.Book{ID:-1, Title:title, Author:author, Status:status}
	
	l.AddBook(book)

}


// accepts book id and passes it to be removed from the library
func removebook(l services.LibraryManager) {
	fmt.Println("\n - - - - - - - - - - - - - - - - ")
	
	var id int

	fmt.Println("Book id: ")
	id = inputInt()
	l.RemoveBook(id)
	

}

// accepts the id of the member and the book to be borrowed 
func borrowbook(l services.LibraryManager) {
	fmt.Println("\n - - - - - - - - - - - - - - - - ")

	var bookid, memberid int
	fmt.Println("Member id: ")
	memberid = inputInt()
	fmt.Println("Book id: ")
	bookid = inputInt()

	l.BorrowBook(bookid, memberid)
	

}

// accepts member id and displays the list of books the user borrowed 
func listborrowedbook(l services.LibraryManager) {
	fmt.Println("\n - - - - - - - - - - - - - - - -")

	var memberid int
	fmt.Println("Member id: ")
	memberid = inputInt()

	fmt.Println()
	if len(l.ListBorrowedBooks(memberid)) == 0 {
		fmt.Println("No Books borrowed")
		return
	}
	for _, j := range l.ListBorrowedBooks(memberid) {
		fmt.Printf("ID: %d , Title: %s\n", j.ID, j.Title)
		fmt.Printf("Author: %s\n", j.Author)
	}
	

}

// lists all books that are available to be borrowed in the library
func listavailablebooks(l services.LibraryManager) {
	fmt.Println("\n - - - - - - - - - - - - - - - - ")

	fmt.Println("Available books: ")
	for _, j := range l.ListAvailableBooks() {
		fmt.Printf("\nID: %d , Title: %s\n", j.ID, j.Title)
		fmt.Printf("Author: %s\n", j.Author)
	}
	

}

//adds a member to the library
func addmember(l services.LibraryManager) {

	var name string
	fmt.Println("Full name:")
	name = input()
	l.AddMember(name)
	

}

// accepts member and book id to be returned to the library
func returnbook(l services.LibraryManager) {
	fmt.Println("\n - - - - - - - - - - - - - - - - ")
	

	var bookid, memberid int
	fmt.Println("Member id: ")
	memberid = inputInt()
	fmt.Println("Book id: ")
	bookid = inputInt()

	l.ReturnBook(bookid, memberid)
	

}


// input handling functions
func input()string{
	buff := bufio.NewReader(os.Stdin)
	userInput,err := buff.ReadString('\n')
	if err != nil {
		log.Fatal("invalid input")
	}
	userInput = strings.TrimSpace(userInput)
	return userInput
}

func inputInt()int{
	buff := bufio.NewReader(os.Stdin)
	userInput,err := buff.ReadString('\n')
	if err != nil {
		log.Fatal("Invalid input \n",err)
	}
	
	userInput = strings.TrimSpace(userInput)
	user,err := strconv.Atoi(userInput) 
	if err != nil {
		log.Fatal("Invalid input \n",err)
	}
	return user
}