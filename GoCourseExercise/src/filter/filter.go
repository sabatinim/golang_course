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
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	log.SetFlags(0)

	basePath := handleCommandLine()
	channel1 := readdir(basePath)
	channel2 := analize_file(".txt", channel1)
	channel3 := apply_regex("[A-Za-z0-9.]+@[a-z]+[a-z.]{2,4}", channel2)
	sink(channel3)
}

func handleCommandLine() string {
	basepath := flag.String("basepath", ".", "insert base path")
	flag.Parse()
	return *basepath
}

func readdir(basePath string) <-chan string {
	out := make(chan string, 1000)
	go func() {

		filepath.Walk(basePath,
			func(path string, info os.FileInfo, err error) error {
				out <- path
				//fmt.Println(path)
				return nil
			})
		close(out)
	}()
	return out
}

func analize_file(extension string, in <-chan string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for path := range in {
			ext := filepath.Ext(path)
			if ext == extension {
				//fmt.Println(ext)
				out <- path
			}
		}
		close(out)
	}()
	return out
}

type Return struct {
	Path  string
	Match string
}

func apply_regex(regex string, in <-chan string) <-chan Return {

	r, _ := regexp.Compile(regex)

	out := make(chan Return, 1000)
	go func() {

		for path := range in {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			reader := bufio.NewReader(file)
			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				matched := r.MatchString(scanner.Text())
				if matched {
					res := r.FindString(scanner.Text())
					out <- Return{Match: res, Path: path}
				}
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		close(out)
	}()
	return out
}

func sink(in <-chan Return) {
	for filename := range in {
		fmt.Println("Found=", filename.Match, " into ", filename.Path)
	}
}
