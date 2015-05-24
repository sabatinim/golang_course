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
