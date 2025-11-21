package models

// defines a user who is a member of the library
type Member struct{
	ID int
	Name string
	BorrowedBooks []Book
}

func NewMember(id int,name string) Member{
	return Member{id,name,make([]Book,0)}
}