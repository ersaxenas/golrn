package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomnum := rand.Intn(100)
	fmt.Println(randomnum~)
	for x := 0; x < 10; x++ {
		fmt.Println("Please guess a no. between 1 and 100: > ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("failed to read input : ", err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("failed to convert to integer: ", err)
		}
		if guess < randomnum {
			fmt.Println("You guessed less. Try again. ")
		} else if guess > randomnum {
			fmt.Println("You guessed more, Try again. ")
		} else {
			fmt.Println("You guessed right. ")
			break
		}
	}

}
