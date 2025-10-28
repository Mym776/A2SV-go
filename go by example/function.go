package main
import(
	"fmt"
)

func Function(){
	fmt.Println(Function2(10))
}
func Function2(a int) int {
	return a*2
}