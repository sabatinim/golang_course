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
	"regexp"
	"strings"
	"bytes"
	"strconv"
)

func Str_compare() {
	var a, b string
	a = "Hello"
	b = "Hello"
	fmt.Println("\"", a, "\" is the same of \"", b, "\"? (", strings.EqualFold(a, b), ")")
}

func CharAt() {
	a := "Hello World"
	fmt.Println("Hello World[3] is (", a[3], ")")
}

func SplitString() {
	a := "Hello#World"
	var splitted []string
	splitted = strings.Split(a, "#")
	fmt.Println("Splitting Hello#World on # char (", splitted[0], " | ", splitted[1], ")")
}

func MatchString() {
	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.
	r, _ := regexp.Compile("p([a-z]+)ch")
	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.
	fmt.Println(r.MatchString("peach"))
	// This finds the match for the regexp.
	fmt.Println(r.FindString("peach punch"))
	// This also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.
	fmt.Println(r.FindStringIndex("peach punch"))
	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// Similarly this will return information about the
	// indexes of matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// These `All` variants are available for the other
	// functions we saw above as well.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.
	fmt.Println(r.Match([]byte("peach")))
	// When creating constants with regular expressions
	// you can use the `MustCompile` variation of
	// `Compile`. A plain `Compile` won't work for
	// constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	// The `regexp` package can also be used to replace
	// subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	// The `Func` variant allows you to transform matched
	// text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

/**
* https://golang.org/pkg/fmt/ (Documentation about format and print variables)
*/
func FormatString(){
	
	var a string
	
	a="111"
	res,err:=strconv.ParseInt(a,2,32)
	fmt.Println("XXX",res)
	
	fmt.Println("XXX",err)
	
	}

