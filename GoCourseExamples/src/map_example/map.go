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

func main() {

	fmt.Println("-- map creation --")
	//literal
	var m = map[string]float64{"1": 2.3, "pi": 3.1415}
	fmt.Println("m = ", m)
	//with make func
	m = make(map[string]float64)
	fmt.Println("m = ", m)
	//assignment
	var m1 map[string]float64
	m = m1
	fmt.Println("m = ", m)

	fmt.Println("\n -- indexing map --")
	m = map[string]float64{"1": 2.3, "pi": 3.1415}
	fmt.Println("m[\"1\"] = ", m["1"])
	//same key
	m["2"] = 4.65
	m["2"] = 20
	fmt.Println("m[\"2\"] = ", m["2"])
	//chek if value is into map

	value, present := m["34"]
	fmt.Println("m[\"43\"]=", value, " is present ", present)
	value, present = m["2"]
	fmt.Println("m[\"2\"]=", value, " is present ", present)
	fmt.Println("m[2] deleting ...")
	delete(m, "2")
	value, present = m["2"]
	fmt.Println("m[\"2\"]=", value, " is present ", present)

	fmt.Println("\n -- for range --")
	for key, value := range m {
		fmt.Printf("key %s -> value %g\n", key, value)
	}
}