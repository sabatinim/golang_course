package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"log"
)

type Message struct {
	Name   string
	Body   string
	Time   int64
	Detail Detail
}

type Detail struct {
	Text   string
	Sender string
}

func main() {

	timestamp := time.Now().Unix()

	file, err := os.Create("data/" + strconv.Itoa(int(timestamp)) + "test.json")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	//js :=json.NewEncoder(file)
	m := Message{"Alice", "Hello", 1294706395881547000, Detail{"XXXX ", "XX"}}
	toWrite, _ := json.Marshal(m)
	fmt.Println(toWrite)
	file.Write(toWrite)
	
	dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            fmt.Println(k)
            //if k != "Name" {
            //    delete(v, k)
            //}
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }

	

}
