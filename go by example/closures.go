package main
import(
	"fmt"
)

func funky() func() {
	i:=0
	return func ()  {
		i++
		fmt.Println(i)
	}
}

func Closure(){
	f := funky()
	f()
	f()
	f()
	f()
	f = funky()
	f()
}