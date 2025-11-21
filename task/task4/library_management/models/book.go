package models

// a struct made to define a book
type Book struct{
	ID int
	Title string
	Author string
	Status string // {"available" or "borrowed"}
}
 