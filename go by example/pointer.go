package main
import(
	"fmt"
)

func zero(a int){
	a = 0
}
func zeroz(a *int){
	*a = 0
}

func Point(){
	x:=1
	fmt.Println(x)
	zero(x)
	fmt.Println(x)
	zeroz(&x)
	fmt.Println(x)
}