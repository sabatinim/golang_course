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
	"regexp"
)

func main() {
	r, _ := regexp.Compile("[A-Za-z0-9.]+@[a-z]+[a-z.]{2,4}")
	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.
	fmt.Println("matching email ", r.MatchString("marco.sabatini@wedjaa.net"))
	fmt.Println("matching email ", r.FindString("La mail di marco sabatini è:marco.sabatini@wedjaa.net"))
}
