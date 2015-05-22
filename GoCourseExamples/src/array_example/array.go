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
	var ar = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Println("ar= ", ar)
	fmt.Println("a= ", a)
	fmt.Println("len(a)= ", len(a), " cap(a)=", cap(a))
	var b = ar[0:4]
	fmt.Println("b= ", b)
	fmt.Println("len(b)= ", len(b), " cap(b)=", cap(b))

	fmt.Println("\n-- resizing  slice --")

	var sl = make([]int, 0, 100)
	fmt.Println("len(sl)= ", len(sl), " cap(sl)=", cap(sl))
	for i := 0; i < 50; i++ {
		len := len(sl)
		sl = sl[0 : i+1]
		sl[len] = i + 1
	}
	fmt.Println("len(sl)= ", len(sl), " cap(sl)=", cap(sl))

	fmt.Println("\n-- append --")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ",slice2)

	fmt.Println("\n-- copy --")
	slice2 = make([]int, 2)
	copy(slice2,slice1)
	fmt.Println("slice1 original = ", slice1)
	fmt.Println("slice2 copy = ",slice2)
	
	fmt.Println("\n-- range --")
	for index,value :=range originalArr{
		fmt.Println("i:[",index,"] v:[",value,"]")
		
		}
	
	

}
