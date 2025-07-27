package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//Define variables
	var name string
	var address string
	var userMap map[string]string

	var scanner = bufio.NewScanner(os.Stdin) // Define and Initialize the scanner to read from standard input
	userMap = make(map[string]string)        // Initialize the map to store user inputs

	// Prompt and scan user input for name
	fmt.Printf("Please enter your name: ")
	name = readUserInputs(scanner)

	// Prompt and scan user input for address
	fmt.Printf("Please enter your address: ")
	address = readUserInputs(scanner)

	// Store the inputs in the map
	userMap["name"] = name
	userMap["address"] = address

	barr, err := json.Marshal(userMap) // Convert the map to JSON format

	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println("JSON Output:", string(barr)) // Print the JSON output
}

func readUserInputs(scanner *bufio.Scanner) string {
	var userInput string

	if scanner.Scan() {
		userInput = scanner.Text()
	}

	return userInput
}
