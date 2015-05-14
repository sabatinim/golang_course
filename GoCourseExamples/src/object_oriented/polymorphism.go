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

package object_oriented

import (
	"fmt"
	"reflect"
	)

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

func Polymorphism() {
   r := Rectangle{length:5, width:3}
   q := Square{side:5}
   shapesArr := [...]Shaper{r, q}

   fmt.Println("Looping through shapes for area ...")
   for n, _ := range shapesArr {
       fmt.Println("Shape details: ", shapesArr[n])
       fmt.Println("Area of this shape is: ", shapesArr[n].Area())
       fmt.Println(reflect.TypeOf(shapesArr[n]))
   }
}
