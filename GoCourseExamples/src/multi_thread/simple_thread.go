package multi_thread

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}

	time.Sleep(time.Second*10)
	fmt.Println("END THREAD")
}

func ExecThread() {
	//go f(0)
	//go f(0)
	for i:=0;i<100;i++{
		go f(0)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("END MAIN THREAD")
}