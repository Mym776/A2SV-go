package main
import(
	"fmt"
)

type book struct{
	name string
	author string
}

func bookFactory(name string) *book{
	var n = book{name, "unknown"}
	return &n
}

func Struc(){
	var hp = book{name:"harryPotter", author:"J.k rowling"}
	var lotr = book{"lord of the rings","tolkein"}
	fmt.Println(hp)
	fmt.Println(lotr)

	k := bookFactory("lotm")
	fmt.Println(k)
}