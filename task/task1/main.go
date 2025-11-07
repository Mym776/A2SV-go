package main

import(
	"fmt"
)

func sum(nums []int)int{
	
	tot :=0

	for _,i := range nums {
		tot +=i
	}
	return tot
}

func main(){
	// test case 1
	var nums = []int{10,20,30}
	sumation := sum(nums)
	fmt.Println("Sum: ",sumation)
	// test case 2  
	nums = []int{}
	sumation = sum(nums)
	fmt.Println("Sum: ",sumation)
	// test case 3
	nums = []int{-1,-2,-3,-4}
	sumation = sum(nums)
	fmt.Println("Sum: ",sumation)
}