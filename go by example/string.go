package main
import(
	"fmt"

)

func Strin(){
	const x = "hello"

	for ind, val:= range x {
		fmt.Printf("%d = %#U is also %v as runes \n", ind,val,val)
	}

	// runes are a bit confusing and unique in go 
	// runes are encased in single quotations
}