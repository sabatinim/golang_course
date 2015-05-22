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
	"unsafe"
)

func main() {

	fmt.Println("--- Integer ---")
	var integer int
	var integer32 int32
	var integer64 int64

	var uinteger uint
	var uintegerptr uintptr

	fmt.Println("SizeOf(integer)", unsafe.Sizeof(integer))
	fmt.Println("SizeOf(integer32)", unsafe.Sizeof(integer32))
	fmt.Println("SizeOf(integer64)", unsafe.Sizeof(integer64))
	fmt.Println("SizeOf(uinteger)", unsafe.Sizeof(uinteger))
	fmt.Println("SizeOf(uintegerptr)", unsafe.Sizeof(uintegerptr), "\n")

	fmt.Println("--- Byte & Rune ---")
	var bytetype byte
	var runetype rune
	fmt.Println("SizeOf(bytetype)", unsafe.Sizeof(bytetype))
	fmt.Println("SizeOf(runetype)", unsafe.Sizeof(runetype), "\n")

	fmt.Println("--- Floating ---")
	var floattype32 float32
	var floattype64 float64
	fmt.Println("SizeOf(floattype32)", unsafe.Sizeof(floattype32))
	fmt.Println("SizeOf(floattype64)", unsafe.Sizeof(floattype64), "\n")

	fmt.Println("--- Complex ---")
	var complextype64 complex64
	var complextype128 complex128
	fmt.Println("SizeOf(complextype64)", unsafe.Sizeof(complextype64))
	fmt.Println("SizeOf(complextype128)", unsafe.Sizeof(complextype128), "\n")

	fmt.Println("--- Operations ---")
	var a, b int = 10, 11
	fmt.Printf("%d + %d = %d\n", a, b, (a + b))
	fmt.Printf("%d - %d = %d\n", a, b, (a - b))
	fmt.Printf("%d * %d = %d\n", a, b, (a * b))
	fmt.Printf("%d / %d = %d\n\n", a, b, (a / b))
	
	var c,d float64 = 2,3
	fmt.Printf("%f + %f = %f\n", c, d, (c + d))
	fmt.Printf("%f - %f = %f\n", c, d, (c - d))
	fmt.Printf("%f * %f = %f\n", c, d, (c * d))
	fmt.Printf("%f / %f = %f\n\n", c, d, (c / d))
	

}
