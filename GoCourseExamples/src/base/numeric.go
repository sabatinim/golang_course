package base

import (
	"fmt"
	"unsafe"
)

func Integer() {

	var a, b int64
	
var c byte
	a = 1
	b = 2
	fmt.Println("SizeOf(int)", unsafe.Sizeof(c) )	
	fmt.Println("a+b=", (a + b) )
	fmt.Println("a-b=", (a - b) )
	fmt.Println("a * b=", (a * b) )
	fmt.Println("a / b=", (a / b) )	
}

func Floating() {

}
