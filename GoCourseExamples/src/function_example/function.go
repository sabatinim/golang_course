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
	"fmt"
)

func f(x int) {
	fmt.Println("x=", x)
}

var t int = 100

func ft() {
	fmt.Println("t=", t)
}

//multiple value return
func multiple_return() (int, int) {
	return 5, 6
}

//variadic func
func variadic_func(args ...int) int {
	tot := 0
	for _, val_arg := range args {
		tot += val_arg
	}
	return tot
}

//closure
func eventgenerator() func() int {
	i := int(0)
	return func() int {
		ret := i
		i += 1
		return ret
	}
}

//defer
func first() {
	fmt.Println("first")
}
func second() {
	fmt.Println("second")
}

//defer & panic

func defer_and_panic() {

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}

//higher order func
func slice_index(limit int, predicate func(val int) bool) int {

	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("--- variable scope ---")
	var x int = 10
	f(x)
	ft()
	fmt.Println("\n--- returning multiple values ---")
	one, two := multiple_return()

	fmt.Printf("one,two := multiple_return() - one=%d two=%d\n", one, two)
	fmt.Println("\n--- variadic function ---")
	fmt.Println("variadic_func(1,2,3,4) = ", variadic_func(1, 2, 3, 4))

	fmt.Println("\n--- closure 1/2 ---")
	add := func(a, b int) int {
		return a + b
	}
	ret := add(1, 3)
	fmt.Println("add(1,3) = ", ret)
	//parameter scope
	var counter int = 0
	increment := func() {
		counter += 1
	}
	fmt.Println("pointer to increment func = ", increment)
	fmt.Println("counter = ", counter, " - execute incrment()")
	increment()
	fmt.Println("counter = ", counter, " - execute incrment()")
	increment()
	fmt.Println("counter = ", counter)

	fmt.Println("\n--- closure 2/2 ---")
	count := eventgenerator()
	fmt.Println("count() ", count())
	fmt.Println("count() ", count())
	fmt.Println("count() ", count())

	fmt.Println("\n--- defer ---")
	defer first()
	second()
	second()

	fmt.Println("\n--- defer&panic ---")
	defer_and_panic()

	fmt.Println("\n--- higher order function ---")
	xs:=[]int{1,2,3,5}
	fmt.Println("-",slice_index(len(xs),func(i int)bool{ return xs[i]==5 }))

}
