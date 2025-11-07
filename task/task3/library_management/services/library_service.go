package services

import (
	"fmt"
	"library/models"
)

type Library struct {
	BookList   map[int]models.Book
	MemberList map[int]models.Member
}

type LibraryManager interface{
	AddBook(models.Book)
	RemoveBook(int)
	BorrowBook(int, int)
	ReturnBook(int, int)
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(int) []models.Book
	AddMember(string)
}

func NewLibrary() Library{
	return Library{make(map[int]models.Book), make(map[int]models.Member)}
}

func (l Library) AddBook(book models.Book) {
	book.ID = len(l.BookList)
	_, exist := l.BookList[book.ID]
	if exist == false {
		l.BookList[book.ID] = book
		fmt.Println(book.ID, book.Title, " added to the Library")
	} else {
		fmt.Println(book.ID, book.Title, " is already in the Library")
	}

}

func (l Library) RemoveBook(bookID int) {
	_, exist := l.BookList[bookID]
	if exist {
		delete(l.BookList, bookID)
		fmt.Println("book removed from Library")
	} else {
		fmt.Println("Book not found")
	}
}

func (l Library) BorrowBook(bookID int, memberID int) {
	borrowedbook, bookExist := l.BookList[bookID]
	member, memberExist := l.MemberList[memberID]
	if !bookExist || !memberExist{
		fmt.Println("Invalid entry")
		return 
	}
	if borrowedbook.Status == "borrowed" {
		fmt.Println("Book not available")
		return
	}
	borrowedbook.Status = "borrowed"
	l.BookList[bookID] = borrowedbook
	member.BorrowedBooks =append(member.BorrowedBooks,borrowedbook)
	l.MemberList[memberID] = member
	fmt.Println(bookID," borrowed to member ", memberID)

}

func (l Library) ReturnBook(bookID int, memberID int) {
	borrowedbook, bookExist := l.BookList[bookID]
	member, memberExist := l.MemberList[memberID]
	if !bookExist || !memberExist{
		fmt.Println("Invalid entry")
		return 
	}
	
	borrowedbook.Status = "borrowed"
	l.BookList[bookID] = borrowedbook
	member.BorrowedBooks =append(member.BorrowedBooks,borrowedbook)
	l.MemberList[memberID] = member
	fmt.Println(bookID," borrowed to member ", memberID)

}

func (l Library) ListAvailableBooks() []models.Book{
	var availableBooks []models.Book
	for _, book := range l.BookList{
		if book.Status=="available"{
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func(l Library) ListBorrowedBooks(memberID int) []models.Book{
	
	m, memberExist := l.MemberList[memberID]
	if !memberExist{
		fmt.Println("member does not exist")
		return nil
	}
	return m.BorrowedBooks
}

func(l Library) AddMember(name string){
	id := len(l.MemberList)
	m := models.NewMember(id, name)
	l.MemberList[id] = m
	fmt.Println(m.Name," add as a member")
}