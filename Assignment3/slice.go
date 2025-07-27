package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Define varibales
	var userInput string

	slice := make([]int, 0, 3)            // Create a slice of integers with length 3
	scanner := bufio.NewScanner(os.Stdin) // Define a scanner to read user input

	for {
		fmt.Printf("Please enter a number: ")
		if scanner.Scan() {
			userInput = scanner.Text()
		}

		// Exit the loop if the user inputs 'X' or 'x'
		if isExitCommand(userInput) {
			fmt.Println("Exiting the program.")
			break
		}

		// Convert userInput to int and handle any conversion errors
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		slice = append(slice, num)           // Append the number to the slice
		sortSlice(slice)                     // Sort the slice after each input
		fmt.Println("Current slice:", slice) // Print the current state of the slice
	}
}

// isExitCommand checks if the provided userInput string is an exit command ("X" or "x").
// It returns true if the input matches either "X" or "x", indicating the user wants to exit.
func isExitCommand(userInput string) bool {
	return userInput == "X" || userInput == "x"
}

// Sort the slice of integers in ascending order
func sortSlice(slice []int) {
	sort.Ints(slice)
}
