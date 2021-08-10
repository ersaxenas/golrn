package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade: ")
	breader := bufio.NewReader(os.Stdin)
	input, err := breader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read input :", err)
	}
	input = strings.TrimSpace(input)
    grades, err := strconv.ParseFloat(input, 64)
    if err != nil {
    	log.Fatal("failed to convert user input to numeric value: ", err)
	}
	var status string
	if grades >= 60 {
       status = "passing"
	} else if grades < 60 {
		status = "failing"
	}
	fmt.Println(status)
}
