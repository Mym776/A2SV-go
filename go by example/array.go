package main
import("fmt")

func Array(){
	
	var a [5]bool
	fmt.Println(a) 
	
	var b  = [5]int{1,2,3,4};
	fmt.Println(b)

	var c = [...]string{"let", "me", "out"}
	fmt.Println(c, len(c))

}	
