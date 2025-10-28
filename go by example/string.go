package main
import(
	"fmt"

)

func Strin(){
	const x = "hello"

	for ind, val:= range x {
		fmt.Printf("%d = %#U \n", ind,val)
	}
}