package main

import (
	"bufio"
	"calculator/internal/handlers"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Calculator application started...")

	for {
		fmt.Println("Input desired operation. Example: \n1 + 2 \nI * VI")
		fmt.Println("Input:")
		if scanner.Scan() {
			input := scanner.Text()

			if strings.ToLower(input) == "exit" {
				fmt.Println("Exiting...")
				break
			}

			answer, err := handlers.Handle(input)

			if err != nil {
				log.Println(err)
			} else {
				fmt.Printf("Answer is %v \n", answer)
			}
		}
	}

	return nil
}
