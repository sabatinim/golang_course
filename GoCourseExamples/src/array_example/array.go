package main

import (
	"fmt"
)

func f(myArr [10]int64) {
	fmt.Println("copy myArr = ", myArr)
}

func fp(myArr *[10]int64) {
	fmt.Println("pointer myArr = ", myArr)

}

func main() {
	fmt.Println("-- array init --")
	var myArray [10]int64
	fmt.Println("myArray[10]=", myArray)

	fmt.Println("\n-- array literals --")
	first := [3]int{1, 2, 3}
	fmt.Println("first := ", first)
	//instance second array 10 element memory only 3 init
	var second = [10]int{1, 2, 3}
	fmt.Println("second = ", second)

	//init only some indexes array values
	var third = [10]int{1: 5, 6: 4, 4: 2}
	fmt.Println("third = ", third)

	fmt.Println("\n-- array are values --")
	f(myArray)
	fp(&myArray)

	fmt.Println("\n-- pointers to array --")
	for i := 0; i < 3; i++ {
		fp(&[10]int64{int64(i), int64(i * i), int64(i * i * i)})
	}

	fmt.Println("\n-- slices --")
	var originalArr [10]int
	var mySlice = originalArr[0:5]
	fmt.Println("originalArr = ", originalArr, " len(originalArr)=", len(originalArr))
	fmt.Println("mySlice = ", mySlice, " len(mySlice)=", len(mySlice))
	fmt.Println("\n-- slice shorthands --")
	var zeroToN = originalArr[:7]
	fmt.Println("originalArr[:7] = ", zeroToN, " len(zeroToN)=", len(zeroToN))
	var fourToN = originalArr[4:]
	fmt.Println("originalArr[4:] = ", fourToN, " len(fourToN)=", len(fourToN))
	var NToN = originalArr[:]
	fmt.Println("originalArr[:] = ", NToN, " len(NToN)=", len(NToN))

	fmt.Println("\n-- slices conceptually --")
	originalArr[1] = 1
	originalArr[3] = 3
	originalArr[8] = 8
	fmt.Println("originalArr init = ", originalArr)

	var fromOrig = originalArr[6:]
	fmt.Println("fromOrig original = ", fromOrig)
	fromOrig[2] = 23
	fmt.Println("fromOrig changed = ", fromOrig)
	fmt.Println("originalArr end = ", originalArr)

	fmt.Println("\n-- make slice --")
	var makeSlice = make([]int, 10)
	fmt.Println("makeSlice  = ", makeSlice)
	fmt.Println("\n-- slice capacity --")


}
