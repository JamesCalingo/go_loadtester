package main

import (
	"challenges/internal/loadtester"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running")
	url := os.Args[1]
	method := os.Args[2]
	status := loadtester.CheckStatus(url, method)
	if status == 200 {
		fmt.Printf("Successfully connected to %s using method %s\n", url, strings.ToUpper(method))
		amount, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		if amount <= 0 {
			log.Fatal("You are attempting to run zero or fewer tests. Please input a positive integer")
		}
		fmt.Println(loadtester.GenerateReport(url, method, amount))
	}
}
