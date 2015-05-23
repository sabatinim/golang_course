package main

import (
	"fmt"
	"math"
)

//struct and function
type MyPoint struct {
	x, y float64
}
func (p *MyPoint) Abs() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

type MyInt int
func (i MyInt) MyRec() int {
	return 4
}

func main() {

	fmt.Println("-- structs --")
	type Point struct {
		x, y float64
	}

	fmt.Println("\n-- making structs --")

	var point Point
	fmt.Println("Point =", point)
	pp := new(Point)
	fmt.Println("pp =", *pp)
	//construct
	point = Point{2, 3}
	fmt.Println("Point{2,3} =", point)
	point = Point{y: 2, x: 3}
	fmt.Println("Point{y:2,x:3} =", point)
	*pp = Point{4, 5}
	fmt.Println("Point{4,5} =", *pp)
	pp = &Point{4, 5}
	fmt.Println("&Point{4,5} =", *pp)

	fmt.Println("\n-- inheritance conflict example --")
	type A struct {
		x int
	}
	type B struct {
		A
		y int
	}
	type C struct {
		A
		B
	}
	var c C
	c.x = 12
	fmt.Println("c= ", c)
	fmt.Println("\n-- struct function --")
	var myPoint = MyPoint{3, 4}
	fmt.Println("myPoint=", myPoint, " myPoint.Abs()=", myPoint.Abs())

}
