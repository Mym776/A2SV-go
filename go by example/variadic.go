package main
import(
	"fmt"
)

func avg(nums ...int) int{
	total:=0
	for _, num:= range nums {
		total += num
		fmt.Println(total)
	}
	
	total /= len(nums)
	return total
}

func variadic(){
	fmt.Println(avg(1,2,3,4,5,6,7,8,9))

	//
	var a = []int{1,2,3,1,2,3,1,2,3,4}
	fmt.Println(avg(a...))
}