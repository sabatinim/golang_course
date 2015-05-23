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

/** Base Object Definition */
type Item struct {
	id       string  // Named field (aggregation)
	price    float64 // Named field (aggregation)
	quantity int     // Named field (aggregation)
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}
func InstanceStruct() {
	item := Item{"xBox_One", 3.56, 5}
	fmt.Println("Instance Address ", item)
	fmt.Println("", item.quantity, " - ", item.id, " cost is ", item.Cost())
}

/**
* Overriding
 */
type SpecialItem struct {
	Item
	catalogId int
}
/* Decomment to specialize implementation
func (sItem *SpecialItem) Cost() float64 {
	return sItem.price * float64(sItem.quantity)*2
}
*/
func main() {
	specialItem := SpecialItem{Item{"xBox_One", 3.56, 5}, 2}
	fmt.Println("Instance Address ", specialItem)
	fmt.Println("", specialItem.quantity, " - ", specialItem.id, " cost is ", specialItem.Cost())
}

/***/