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
	
	//read from stream and write json encoded
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
        }
        if err := enc.Encode(&v); err != nil {
            log.Println(err)
        }
    }
}
