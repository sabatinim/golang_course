package main

import (
	"fmt"
	"db_init"
)

func main() {

	myDb := db_init.GetDb();
	fmt.Print("DB ",myDb)
}
