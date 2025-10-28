package main
import(
	"fmt"
)

func Maps(){
	var maps = make(map[string]int)
	maps["hello"]=121
	fmt.Println(maps["hello"])
	//blank checker
	var _, blank = maps["hola"]
	fmt.Println(blank)
}