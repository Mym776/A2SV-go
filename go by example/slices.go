package main
import(
	"fmt"
)

func Slices(){
	var s = make([]int,3)
	fmt.Println(s)
	
	t:= s[0:1]
	fmt.Println(t)
	t = append(s, 12)
	fmt.Println(t)

}