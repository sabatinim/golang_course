package main

import (
	"fmt"
	//"time"
)

func pump(channel chan int) {

	for i := 0; ; i++ {
		fmt.Println("Sending :", i)
		channel <- i
	}
}

func suck(id int,channel chan int) {
	
	for {
		//time.Sleep(time.Second*1)
		fmt.Println("Sucker", id," get:", <-channel)
		}
}

func pumpWithChann() chan int{
	
		var intChannel = make(chan int)
		go func(){
			for i:=0;;i++{
				intChannel <-i 
				}
			}()
		return intChannel
	}

func main() {

	//var intChannel = make(chan int)
	//go pump(intChannel)
	//fmt.Println(<-intChannel)
	//go suck(intChannel)
	
	stream:=pumpWithChann()
	go suck(1,stream)
	go suck(2,stream)
	
	var input string
	fmt.Scan(&input)
	
}
