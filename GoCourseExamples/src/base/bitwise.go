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

package base

import (
	"fmt"
)

func Bitwise() {
	// Use bitwise OR | to get the bits that are in 1 OR 2
	// 1     = 00000001
	// 2     = 00000010
	// 1 | 2 = 00000011 = 3
	fmt.Println(1 | 2)

	// Use bitwise OR | to get the bits that are in 1 OR 5
	// 1     = 00000001
	// 5     = 00000101
	// 1 | 5 = 00000101 = 5
	fmt.Println(1 | 5)

	// Use bitwise XOR ^ to get the bits that are in 3 OR 6 BUT NOT BOTH
	// 3     = 00000011
	// 6     = 00000110
	// 3 ^ 6 = 00000101 = 5
	fmt.Println(3 ^ 6)

	// Use bitwise AND & to get the bits that are in 3 AND 6
	// 3     = 00000011
	// 6     = 00000110
	// 3 & 6 = 00000010 = 2
	fmt.Println(3 & 6)

	// Use bit clear AND NOT &^ to get the bits that are in 3 AND NOT 6 (order matters)
	// 3      = 00000011
	// 6      = 00000110
	// 3 &^ 6 = 00000001 = 1
	fmt.Println(3 &^ 6)
}