package services

import (
	"fmt"
	"library/models"
	"time"
)

// library that stores members and books
type Library struct {
	BookList   map[int]models.Book
	MemberList map[int]models.Member
}

type LibraryManager interface {
	AddBook(models.Book)
	RemoveBook(int)
	BorrowBook(int, int)
	ReturnBook(int, int)
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(int) []models.Book
	AddMember(string)
	ReserveBook(int, int)
	UnReserveBook(int,int)
}

//returns a new instance of a library
func NewLibrary() Library {
	return Library{make(map[int]models.Book), make(map[int]models.Member)}
}

// adds a book to the library
func (l Library) AddBook(book models.Book) {
	book.ID = len(l.BookList)
	_, exist := l.BookList[book.ID]
	if !exist {
		l.BookList[book.ID] = book
		fmt.Println("\n",book.Title, " by ", book.Author," added to the Library")
	} else {
		fmt.Println(book.Title, " is already in the Library")
	}

}

//removes a book from the library
func (l Library) RemoveBook(bookID int) {
	removed, exist := l.BookList[bookID]
	if exist && removed.Status=="available"{
		delete(l.BookList, bookID)
		fmt.Println(removed.Title," removed from Library")
	} else {
		fmt.Println("Book not found")
	}
}

// takes a book and gives it to a member 
func (l Library) BorrowBook(bookID int, memberID int) {
	borrowedbook, bookExist := l.BookList[bookID]
	member, memberExist := l.MemberList[memberID]
	if !bookExist || !memberExist {
		fmt.Println("Invalid entry")
		return
	}
	if borrowedbook.Status == "borrowed" {
		fmt.Println("Book not available")
		return
	}else if borrowedbook.Status == "reserved"{
		reserved := false
		for i,j := range member.BorrowedBooks{
			if j == borrowedbook {
				fmt.Println("reserved book successfuly borrowed")
				reserved = true
				member.BorrowedBooks = append(member.BorrowedBooks[:i],member.BorrowedBooks[i+1:]... )
				break
			}
		}
		if !reserved{
			fmt.Println("Book not available")
			return 
		}else{
		}
	}
	borrowedbook.Status = "borrowed"
	l.BookList[bookID] = borrowedbook
	member.BorrowedBooks = append(member.BorrowedBooks, borrowedbook)
	l.MemberList[memberID] = member
	fmt.Println(borrowedbook.Title, " borrowed to member ", member.Name)

}

// accepts a book returned by a member
func (l Library) ReturnBook(bookID int, memberID int) {

	member, memberExist := l.MemberList[memberID]
	borrowedbook, bookExist := l.BookList[bookID]
	if !bookExist || !memberExist {
		fmt.Println("Invalid entry")
		return
	}
	var memberBook bool = false

	for i, j := range member.BorrowedBooks {
		if j == borrowedbook {
			memberBook = true

			//remove the book from the member's borrowed book list
			if len(member.BorrowedBooks) == 1 {
				member.BorrowedBooks = []models.Book{}
			} else if len(member.BorrowedBooks) == i+1 {
				member.BorrowedBooks = member.BorrowedBooks[:len(member.BorrowedBooks)-1]
			} else {
				member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1])

			}
			break
		}
	}
	if !memberBook {
		fmt.Println("Book was not borrowed by memeber : ", member.Name)
		return
	} else {
		borrowedbook.Status = "available"
		l.BookList[bookID]=borrowedbook
		l.MemberList[memberID]=member
		fmt.Println(borrowedbook.Title, " successfully returned to the library")
		
	}

}

// lists all books available to be borrowed 
func (l Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.BookList {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// lists all books borrowed by a specific member
func (l Library) ListBorrowedBooks(memberID int) []models.Book {

	m, memberExist := l.MemberList[memberID]
	if !memberExist {
		fmt.Println("member does not exist")
		return nil
	}
	return m.BorrowedBooks
}

// adds a new member to the library
func (l Library) AddMember(name string) {
	id := len(l.MemberList)
	m := models.NewMember(id, name)
	l.MemberList[id] = m
	fmt.Println(m.Name, " added as a member, ID: ",m.ID)
}


func(l Library) ReserveBook(memberID int, bookID int) {
	
	book, bookExist := l.BookList[bookID]
	member, memberExist := l.MemberList[memberID]

	if !bookExist {
		fmt.Println("book does not exist")
		return
	}

	if !memberExist {
		fmt.Println("member does not exist")
		return
	}

	
	if book.Status != "available"{
		fmt.Println("Book not available for reservation")
		return 
	} 
		
	book.Status = "reserved"
	member.BorrowedBooks = append(member.BorrowedBooks, book)

	l.BookList[bookID] = book
	l.MemberList[memberID] = member
	fmt.Println("book reserved, valid for 15 seconds")
	
}

func(l Library) UnReserveBook(memberID int, bookID int) {
	time.Sleep(15*time.Second)
	book, bookExist := l.BookList[bookID]
	_, memberExist := l.MemberList[memberID]

	if !bookExist {
		fmt.Println("book does not exist")
		return
	}

	if !memberExist {
		fmt.Println("member does not exist")
		return
	}

	
	if book.Status == "borrowed"{
		return 
	} 
	
	l.ReturnBook(bookID,memberID)
	fmt.Println("book unreserved")
}

