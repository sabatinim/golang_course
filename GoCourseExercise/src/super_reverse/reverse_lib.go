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

package super_reverse

import (
	"unicode/utf8"
)

// IsPalindrome in ascii
func IsPalindromeAscii(s string) bool {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return IsPalindromeAscii(s[1 : len(s)-1])
}

// IsPalindrome in utf 8 version
func IsPalindromeUtf8(s string) bool {
	if utf8.RuneCountInString(s) <= 1 {
		return true
	}
	first, sizeOfFirst := utf8.DecodeRuneInString(s)
	last, sizeOfLast := utf8.DecodeLastRuneInString(s)
	if first != last {
		return false
	}
	return IsPalindromeUtf8(s[sizeOfFirst : len(s)-sizeOfLast])
}
