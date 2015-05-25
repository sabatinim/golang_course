/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
)

//interface
type AbsInterface interface {
	Abs() float64
}

type MyFloat float64

func (p MyFloat) Abs() float64 {
	if p < 0.0 {
		return float64(-p)
	}
	return float64(p)
}

type MyPoint struct {
	x, y float64
}

func (p *MyPoint) Abs() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

//polimo
type Shaper interface {
	Area() int
}

type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

type Square struct {
	side int
}

func (sq Square) Area() int {
	return sq.side * sq.side
}


func main() {
	
	fmt.Println("-- interface -- ")

	var intf AbsInterface
	
	var p = MyPoint{3, 4}
	intf = &p
	fmt.Println("ref ", intf.Abs())

	var myFloat MyFloat = -56
	intf = &myFloat
	fmt.Println("ref ", intf.Abs())

	fmt.Println("\n-- writer interface -- ")
	var myWriter = bufio.NewWriter(os.Stdout)
	myWriter.WriteString("Buffered string to STDOUT")
	myWriter.Flush()

	fmt.Println("\n-- polimorphism -- ")	
	r := Rectangle{length: 5, width: 3}
	q := Square{side: 5}
	shapesArr := [...]Shaper{r, q}

	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapesArr {
		fmt.Println("Shape details: ", shapesArr[n])
		fmt.Println("Area of this shape is: ", shapesArr[n].Area())
		fmt.Println(reflect.TypeOf(shapesArr[n]))
	}
	
}
