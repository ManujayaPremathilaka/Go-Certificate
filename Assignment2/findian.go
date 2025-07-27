package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	// Define variables
	var userInput string

	// Prompt and scan user input
	fmt.Printf("Please enter a string: ")
	// Use bufio.Scanner to read the full line including spaces
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		userInput = scanner.Text()
	}

	found := isFound(userInput) // Check if the input meets the criteria

	if found {
		fmt.Printf("Found")
	} else {
		fmt.Printf("Not Found")
	}
}

// isFound checks if the user input contains 'a', starts with 'i', and ends with 'n'
func isFound(userInput string) bool {
	var startsWithI bool = strings.HasPrefix(userInput, "i") || strings.HasPrefix(userInput, "I")
	var hasA bool = strings.Contains(userInput, "a") || strings.Contains(userInput, "A")
	var endsWithN bool = strings.HasSuffix(userInput, "n") || strings.HasSuffix(userInput, "N")

	return startsWithI && hasA && endsWithN
}
