package models

type Member struct{
	ID int
	Name string
	BorrowedBooks []Book
}

func NewMember(id int,name string) Member{
	return Member{id,name,make([]Book,0)}
}