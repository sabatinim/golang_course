package main

import (
	"fmt"
	"os"
	"strings"
)

func first_example() {
	
	who:="World!!!"
	if len(os.Args)>1{
		who=strings.Join(os.Args[1:],"")
		}
	fmt.Println("Hello",who)
}