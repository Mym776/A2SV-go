package main
import(
	"fmt"
)

func Range(){
	var x= []int{9,1,1,2,3,4,5,5}
	for ind, val := range x {
		fmt.Println(fmt.Sprint(ind,"->",val))
	}
}