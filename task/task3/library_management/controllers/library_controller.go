package controllers

import (
	"bufio"
	"fmt"
	"library/models"
	"library/services"
	"os"
	// "strconv"
	"strings"
)

func Start(l services.LibraryManager) {

	menu()

	for {
		var inp int
		fmt.Println("Please input from 1-7: ")
		fmt.Scan(&inp)
		
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
			fmt.Println(inp," is not a valid choice")

		}
	}
	
}

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

func addbook(l services.LibraryManager) {
	fmt.Println("\n - - - - - - - - - - - - - - - - ")

	reader := bufio.NewReader(os.Stdin)
	
	var title string
	var author string
	var status string = "available"

	
	
	fmt.Println("Book Title")
	title, _ = reader.ReadString('\n')

	fmt.Println("Book author")
	author, _ = reader.ReadString('\n')

	title = strings.TrimSpace(title)
	author = strings.TrimSpace(author)

	var book = models.Book{-1, title, author, status}
	l.AddBook(book)
	listavailablebooks(l)
}

func removebook(l services.LibraryManager) {
	fmt.Println("\n - - - - - - - - - - - - - - - - ")
	
	var id int
	fmt.Println("Book id: ")
	fmt.Scan(&id)
	l.RemoveBook(id)

}

func borrowbook(l services.LibraryManager){
	fmt.Println("\n - - - - - - - - - - - - - - - - ")
	
	var bookid, memberid int;
	fmt.Println("Member id: ")
	fmt.Scan(&memberid)
	fmt.Println("Book id: ")
	fmt.Scan(&bookid)

	l.BorrowBook(bookid,memberid)
}

func listborrowedbook(l services.LibraryManager){
	fmt.Println("\n - - - - - - - - - - - - - - - -")
	
	var memberid int;
	fmt.Println("Member id: ")
	fmt.Scan(&memberid)


	fmt.Println()
	if len(l.ListBorrowedBooks(memberid))==0{
		fmt.Println("No Books borrowed")
		return
	}
	for _,j := range l.ListBorrowedBooks(memberid){
		fmt.Printf("ID: %d , Title: %s\n", j.ID,j.Title)
		fmt.Printf("Author: %s\n", j.Author)
	}
}


func listavailablebooks(l services.LibraryManager){
	fmt.Println("\n - - - - - - - - - - - - - - - - ")
	
	fmt.Println("Available books: ")
	for _,j := range l.ListAvailableBooks(){
		fmt.Printf("\nID: %d , Title: %s\n", j.ID,j.Title)
		fmt.Printf("Author: %s\n", j.Author)
	}
}


func addmember(l services.LibraryManager){
	
	var name string;
	fmt.Println("Full name:")
	fmt.Scan(&name)
	l.AddMember(name)

}


func returnbook(l services.LibraryManager){
	fmt.Println("\n - - - - - - - - - - - - - - - - ")
	
	var bookid, memberid int;
	fmt.Println("Member id: ")
	fmt.Scan(&memberid)
	fmt.Println("Book id: ")
	fmt.Scan(&bookid)

	l.ReturnBook(bookid,memberid)
}