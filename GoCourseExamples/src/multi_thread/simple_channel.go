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
