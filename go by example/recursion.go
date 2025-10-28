package main
import(
	"fmt"
)
var mp = make(map[int]int)
func fib(a int) int {
	if a<=2{
		return 1
	}
	if mp[a]>0{
		return mp[a]
	}
	b :=  fib(a-1)
	c:= fib(a-2)
	
	var f = b+c
	mp[a]=f
	fmt.Println(f)
	return f
}

func Rec(){
	fmt.Print(fib(50))
}