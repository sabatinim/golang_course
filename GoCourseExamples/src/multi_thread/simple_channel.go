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

package multi_thread

import (
	"fmt"
	"strconv"
	"time"
)

var i int

func makeCakeAndSend(cs chan string) {
	i = i + 1
	cakeName := "Strawberry Cake " + strconv.Itoa(i)
	fmt.Println("Making a cake and sending ...", cakeName)
	cs <- cakeName //send a strawberry cake
}

func receiveCakeAndPack(cs chan string) {
	s := <-cs //get whatever cake is on the channel
	fmt.Println("Packing received cake: ", s)
}

func ExecChan() {
	cs := make(chan string)
	for i := 0; i < 3; i++ {
		go makeCakeAndSend(cs)
		go receiveCakeAndPack(cs)

		time.Sleep(time.Second * 5)//sleep for 5 secs
	}
	//sleep for a while so that the program doesn’t exit immediately and output is clear for illustration
	var input string
	fmt.Scanln(&input)

}
