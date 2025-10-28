package main

import("fmt")

func Switch(){
	x:="hello"
	switch x {
	case "hallo":
		fmt.Println("hallo")
	case "hello":
		fmt.Println("hello")
	}
}