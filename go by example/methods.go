package main
import(
	"fmt"
)

func (n book) meth1() {
	fmt.Println("book name: ",n.name)
	fmt.Println("book author: ",n.author)
}

func Meth(){
	bk := book{"harry potter","jk rowling"}
	bk.meth1()
}