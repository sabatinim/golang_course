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
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func openAndReadFile() {
	file, err := os.Open("data/test.txt")
	if err != nil {
		// handle the error here
		fmt.Println(err)
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Println(stat)
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func openAndReadFileShort() {
	
	bs, err := ioutil.ReadFile("data/test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func createFile() {
	timestamp := time.Now().Unix()
	file, err := os.Create("data/" + strconv.Itoa(int(timestamp)) + "test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString("test")
}

func readdir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func walkDir() {

	filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})
}

func main() {
	openAndReadFile()
	//openAndReadFileShort()
	//createFile()
	//readdir()
	walkDir()
}
