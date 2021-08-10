package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("File stat program: ")
	//fileInfo, _ := os.Stat("/Users/saurabh.saxena/learnanddev/golang/headfirstgo/src/com/lrn.test/ch2/FileInfo.go")
	fileInfo, err := os.Stat("FileInfo.go")
	if err != nil {
		log.Fatal("Error occurred while reading file stats", err)
	}
	fmt.Println("File name is: " + fileInfo.Name())
}
