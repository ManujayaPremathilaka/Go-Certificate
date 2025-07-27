package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Define variables
	var userInputArr [10]int

	// Get user inputs
	getUserInputs(&userInputArr)
	bubbleSort(userInputArr[:])

	// Print the sorted array
	fmt.Println("Sorted array:", userInputArr)
}

func bubbleSort(userInputSlice []int) {
	for i := 0; i < (len(userInputSlice) - 1); i++ {
		if userInputSlice[i] > userInputSlice[i+1] {
			swap(userInputSlice, i) // Swap the elements if the current element is greater than the next
			i = -1                  // Reset i to -1 to start over after a swap
			continue                // Continue to the next iteration of the outer loop
		}
	}
}

func swap(swapSlice []int, index int) {
	var indexValue int = swapSlice[index] // Store the value at the current index
	swapSlice[index] = swapSlice[index+1] // Swap the current index with the next index
	swapSlice[index+1] = indexValue       // Assign the stored value to the next index
}

func getUserInputs(userInputArr *[10]int) {
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner to read user input

	for i := 0; i < len(userInputArr); i++ {
		fmt.Printf("Please enter number %d: ", (i + 1))

		if scanner.Scan() {
			var input int
			_, err := fmt.Sscanf(scanner.Text(), "%d", &input)

			if err != nil {
				fmt.Println("Invalid input, please enter a valid number.")
				i-- // Decrement i to repeat this iteration
				continue
			}

			userInputArr[i] = input // Store the input in the array
		}
	}
}
