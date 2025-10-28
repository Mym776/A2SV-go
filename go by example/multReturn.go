package main

import "fmt"

func Plus(a int, b int) (string, int) {
	c := fmt.Sprint(a," + ",b," =")
	return c, a + b
}
func p() {
	a, b := Plus(10, 20)
	fmt.Println(a,b)
}