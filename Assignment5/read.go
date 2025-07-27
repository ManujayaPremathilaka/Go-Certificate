package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name struct with two fields
type Name struct {
	fname string
	lname string
}

func main() {
	// Define variables
	var filename string
	var names []Name

	// Prompt for filename
	fmt.Print("Enter the name of the text file: ")
	fmt.Scanln(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner and a slice to store the names
	scanner := bufio.NewScanner(file)

	// Read each line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			first := truncateTo20(parts[0])
			last := truncateTo20(parts[1])
			names = append(names, Name{fname: first, lname: last})
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print all names
	fmt.Println("\nNames read from file:")
	for _, n := range names {
		fmt.Printf("First Name: %-20s Last Name: %-20s\n", n.fname, n.lname)
	}
}

// truncateTo20 ensures the string is no longer than 20 characters
func truncateTo20(s string) string {
	if len(s) > 20 {
		return s[:20]
	}
	return s
}
