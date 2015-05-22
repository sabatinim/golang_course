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

import ("fmt")

func f(x int) {
	x = 0;
}

func f_ptr(x *int) {
	*x = 0;
}


func main() {
	
	fmt.Println("-- pointers overwiew -- ")
	
	var myInteger int = 23
	fmt.Println("myInteger = ",myInteger)
	fmt.Println("f(myInteger)")
	f(myInteger)
	fmt.Println("f(myInteger)")
	fmt.Println("myInteger = ",myInteger)
	f_ptr(&myInteger)
	fmt.Println("f_ptr(myInteger)")
	fmt.Println("myInteger = ",myInteger)
	
	fmt.Println("\n-- * and & operators -- ")
	var myIntegerPointer = &myInteger
	fmt.Println("myInteger value = ",myInteger)
	fmt.Println("myInteger address = ",myIntegerPointer)
	fmt.Println("myIntegerPointer value = ",*myIntegerPointer)
	
	fmt.Println("\n-- new instance -- ")
	var newIntegerPtr = new(int)
	
	*newIntegerPtr = 10
	fmt.Println("newIntegerPtr value = ",*newIntegerPtr)
	f_ptr(newIntegerPtr)
	fmt.Println("newIntegerPtr value = ",*newIntegerPtr)
	
	
	
	
	
	
}
