package main

import (
	"fmt"
	"strconv"
)
const NUM_ELEM = 10000



func doFirstJob(words []string) <-chan string {

	jobs := make(chan string)
	go func() {
		for _, word := range words { // Blocks waiting for a send
			fmt.Println(word) // Do one job
			jobs <- word
		}
		close(jobs)
	}()
	return jobs
}

func doSecondJob(word string, in <-chan string) <-chan bool {
	out := make(chan bool, cap(in))
	go func() {
		for word_send := range in {

			if word == word_send {
				fmt.Println("Found=",word)
			}
			fmt.Println("Check = ",word_send)
			
			out <- true
		}
		close(out)
	}()
	return out
}

func sink(in <-chan bool) {
	for channelWrite := range in {
		fmt.Println(channelWrite)
	}
}

func main() {

	words := make([]string,0)
	for i := 0; i < NUM_ELEM; i++ {
		words = append(words, "Word_"+strconv.Itoa(i))
	}
	chanFirstJob := doFirstJob(words)
	chanSecondJob := doSecondJob("Word_78", chanFirstJob)
	sink(chanSecondJob)
}
