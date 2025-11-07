package main
import(
	"fmt"
)

const pi float32 = 3.1415

type shape interface{
	area()

}


type circle struct{
	radius int
}
type rect struct{
	width, height int
}

func (c circle)area() {

	ar:= float32(c.radius)* float32(c.radius) * pi
	fmt.Println(ar)

}

func measure(s shape){
	s.area()
}

func (r rect) area(){
	ar:= r.height*r.width
	fmt.Println(ar)
}

func Interf(){
	rec := rect{10,20}
	measure(rec)
	cir := circle{10}
	measure(cir)
}