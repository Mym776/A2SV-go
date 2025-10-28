package main
import("fmt")

func If(){
	x:= 987
	fmt.Println("divisible by: ")
	if x%2==0{
		fmt.Println("two")
	}
	if x%3==0{
		fmt.Println("three")
	}
	if x%5==0{
		fmt.Println("five")
	}
	if x%7==0{
		fmt.Println("seven")
	}
	if x%11==0{
		fmt.Println("eleven")
	}
	if x%2==13{
		fmt.Println("thirteen")
	}
	if x%2==17{
		fmt.Println("seventeen")
	}
}