package main

import (
	"fmt"
)

// GenDisplaceFn returns a function that calculates displacement
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64

	// Prompt user for inputs
	fmt.Print("Enter acceleration (a): ")
	fmt.Scanln(&a)

	fmt.Print("Enter initial velocity (v0): ")
	fmt.Scanln(&v0)

	fmt.Print("Enter initial displacement (s0): ")
	fmt.Scanln(&s0)

	// Generate displacement function
	fn := GenDisplaceFn(a, v0, s0)

	fmt.Print("Enter time (t): ")
	fmt.Scanln(&t)

	// Compute and print displacement
	displacement := fn(t)
	fmt.Printf("Displacement after %.2f seconds is: %.2f\n", t, displacement)
}
