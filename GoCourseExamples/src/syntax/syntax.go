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

//declaration
func sum(a, b int) int {
	return a + b
}

func main() {

	//syntax overview 1/2
	fmt.Println("-- syntax overview 1/2 --")
	var a int
	var b, c *int
	var d []int
	fmt.Println("a=",a,"b=", b,"c=", c,"d=", d,"\n") //0 <nil> <nil> []
	
	//syntax overview 2/2
	fmt.Println("-- syntax overview 2/2 --")
	const three = 3
	var i int = three
	fmt.Printf("(three=%d) (i=three) -> ( i=%d )\n",three, i)

	//declaration
	fmt.Println("-- declaration --")
	var declared_var int
	const PI = 22. / 7. //is the same of 22.0/7.0
	var p, q int
	fmt.Println("declared_var=", declared_var, " PI=", PI, " p=", p, " q=", q, " sum(1,1)=", sum(1, 1)) //

	//var identifier
	fmt.Println("-- identifier --")
	var j = 365.245
	var k int = 0
	var l, m uint64 = 3, 4
	var nanosecs int64 = 1e9
	var inter, floater, stringer = 1, 2.0, "string"
	fmt.Println("j=", j, " k=", k, " l=", l, " m=", m, " nanosecs=", nanosecs, " integer=", inter, " floater=", floater, " stringer=", stringer,"\n")

	//distributing var
	fmt.Println("-- distributing var --")
	var (
		variable_scope_int int
	)
	fmt.Println("variable_scope_int=", variable_scope_int,"\n")

	//:= short declaration
	fmt.Println("-- := short declaration --")	
	a_int := 123 //integer
	a_string := "String value"
	a_float := 123.34
	fmt.Println("a_int=", a_int, " a_string=", a_string, " a_float_=", a_float,"\n")

	//const
	fmt.Println("-- const --")	
	const (
		Monday, Thesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	fmt.Println("Monday =", Monday, " Thesday =", Thesday, " Wednesday =", Wednesday)
	fmt.Println("Thursday =", Thursday, " Friday =", Friday, " Saturday =", Saturday,"\n")

	//iota
	fmt.Println("-- iota --")	
	const (
		Monday_iota    = iota
		Thesday_iota   = iota
		Wednesday_iota = iota
		Thursday_iota  = iota
		Friday_iota    = iota
		Saturday_iota  = iota
	)
	fmt.Println("Monday =", Monday_iota, " Thesday =", Thesday_iota, " Wednesday =", Wednesday_iota)
	fmt.Println("Thursday =", Thursday_iota, " Friday =", Friday_iota, " Saturday =", Saturday_iota,"\n")
	const (
		loc = iota
	)//TODO:finire

	//type
	fmt.Println("-- type --")	
	type Point struct {
		x, y, z float64
		name    string
	}
	fmt.Println("Point =", Point{1, 2, 3, " point name"})
	type MySuperInteger int
	var superInt MySuperInteger
	superInt = 23
	fmt.Println("superInt =", superInt,"\n")

	//new
	fmt.Println("-- new --")
	var point *Point = new(Point)
	point.x = 1
	point.y = 1
	point.z = 1
	point.name = "check point charlie"
	fmt.Println("Point =", point,"\n")

	//assignment
	fmt.Println("-- assignment --")
	var sum1, sum2, sum3 = sum(1, 2), sum(3, 4), sum(5, 6)
	fmt.Println("sum1 =", sum1, "sum2 =", sum2, "sum3 =", sum3,"\n")

	//if
	fmt.Println("-- if --")
	var x int
	if x < 5 {
		fmt.Println("LESS")
	}
	x = 5
	if x < 5 {
		fmt.Println("LESS")
	} else if x == 5 {
		fmt.Println("SAME")
	}

	if v := 4; v < 10 {
		fmt.Printf("%d less than 10\n", v)
	} else {
		fmt.Printf("%d not less than 10\n", v)
	}

	//for - while
	fmt.Printf("-- for - loop --")
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d ", i)
	}
	fmt.Printf("\n")

	var brek_forever_loop_it int = 0
	for {
		if brek_forever_loop_it == 20 {
			break
		}
		brek_forever_loop_it += 1
		fmt.Printf("for ;; exec -")
	}
	fmt.Printf("\n")
	
	brek_forever_loop_it = 0
	for {
		if brek_forever_loop_it == 20 {
			break
		}
		brek_forever_loop_it += 1
		fmt.Printf("for exec -")
	}
	fmt.Printf("\n")

	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf(" (i=%d j=%d) ", i, j)
	}
	fmt.Printf("\n")

	//switch
	fmt.Println("-- switch --")
	a = 125
	switch a {
	case 0:
		fmt.Printf("case 0\n")
	default:
		fmt.Printf("default case\n")
	}

	XY, XZ := 5, 6
	switch {
	case XY < XZ:
		fmt.Printf("case %d < %d \n",XY,XZ)
	case XY == XZ:
		fmt.Printf("case %d == %d \n",XY,XZ)
	case XY > XZ:
		fmt.Printf("case %d > %d \n",XY,XZ)
	}
	
	switch XY,XZ = 6,5; {
	case XY < XZ:
		fmt.Printf("case %d < %d \n",XY,XZ)	
		case XY == XZ:
		fmt.Printf("case %d == %d \n",XY,XZ)
	case XY > XZ:
		fmt.Printf("case %d > %d \n",XY,XZ)
	}

}
