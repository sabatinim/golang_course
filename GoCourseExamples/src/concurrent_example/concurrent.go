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
	"time"
)

func pump(channel chan int) {

	for i := 0; ; i++ {
		fmt.Println("Sending :", i)
		channel <- i
	}
}

func suck(id int, channel chan int) {

	for {

		fmt.Println("Sucker", id, " get:", <-channel, "\n")
		time.Sleep(time.Second * 2)
	}
}

func simpleWork() {
	for {
		fmt.Println("Work")
		//time.Sleep(time.Second*5)
	}
}

func pumpWithChann() chan int {
	var intChannel = make(chan int)
	go func() {
		for i := 0; ; i++ {
			fmt.Println("pumpWithChann send = ", i, "")
			intChannel <- i
		}
	}()
	return intChannel
}
func suchWithChann(channel chan int) {
	go func() {

		for v := range channel {
			fmt.Println(v)
		}

	}()
}

func work_routine() (chan bool, chan string) {

	var donechannel = make(chan bool, 10)
	var errchannel = make(chan string, 10)
	go func() {
		i := 0
		for {
			fmt.Println("work_routine Work")
			if i == 0 {
				donechannel <- true
				time.Sleep(time.Second * 1)
				i = 1
			} else {
				errchannel <- "Error on working goroutine"
				i = 0
				time.Sleep(time.Second * 1)

			}
		}
	}()
	return donechannel, errchannel
}

func log_routine(donechannel chan bool, errchannel chan string) {

	for {
		select {
		case v := <-donechannel:
			fmt.Println("log_routine read ", v)
		case v := <-errchannel:
			fmt.Println("log_routine read ", v)

		case v:=<-time.After(time.Second * 1):
			fmt.Println("Time out ",v)

		}
	}

}

func terminate(done chan int) {
	time.Sleep(time.Second * 10)
	done <- 0
}

func main() {

	//var intChannel = make(chan int)
	//go pump(intChannel)
	//fmt.Println(<-intChannel)
	//go suck(intChannel)

	//pump and suck with stream in function pump
	//stream:=pumpWithChann()
	//go suck(1,stream)
	//go suck(2,stream)

	//for i := 0; i < 1000; i++ {
	//	go simpleWork()
	//}
	//pump and suck with stream in function pump and suck
	//suchWithChann(pumpWithChann())
	donechannel, errchannel := work_routine()
	go log_routine(donechannel, errchannel)
	//var input string
	//fmt.Scan(&input)

	var done = make(chan int)
	go terminate(done)
	fmt.Println("Done arrived ", <-done)

}
