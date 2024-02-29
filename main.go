package main

import (
	"challenges/internal/loadtester"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	amount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if amount <= 0 {
		log.Fatal("You are attempting to run zero or fewer tests. Please input a positive integer")
	}
	fmt.Println(loadtester.CheckStatus(os.Args[1], amount))
}
