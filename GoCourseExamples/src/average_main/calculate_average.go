package main

import (
	avg "average_pkg" //alias
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	var returnValue = avg.Average(xs)
	fmt.Println("Average=", returnValue)
}