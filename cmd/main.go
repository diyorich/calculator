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
		fmt.Println("Input desired operation (Example: \n1 + 2 \n10 * 35 ...")
		fmt.Println("Input:")
		if scanner.Scan() {
			input := scanner.Text()

			if strings.ToLower(input) == "exit" {
				fmt.Println("Exiting...")
				break
			}

			err := handlers.Handle(input)

			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}
