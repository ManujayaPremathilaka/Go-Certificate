package main

import (
	"bufio"
	"fmt"
	"os"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	// Initialize animals map
	animals := initAnimals()

	// Define scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	// Loop for user input
	for {
		fmt.Print(">")
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				fmt.Println("Exiting the program.")
				break
			}

			var animalName, action string
			_, err := fmt.Sscan(input, &animalName, &action)
			if err != nil {
				fmt.Println("Invalid input. Please enter in the format: <animal> <action>")
				continue
			}

			var animal *Animal = handleAnimal(animalName, animals)
			if animal == nil {
				continue // Skip to the next iteration if animal is not found
			}
			handleAction(animal, action, animalName)
		}

	}
}

func handleAction(animal *Animal, action string, animalName string) {
	switch action {
	case "eat":
		animal.Eat(animalName)
	case "move":
		animal.Move(animalName)
	case "speak":
		animal.Speak(animalName)
	default:
		fmt.Println("Unknown action. Please use 'eat', 'move', or 'speak'.")
	}
}

func (animal *Animal) Eat(name string) {
	fmt.Printf("The %s eats %s.\n", name, animal.food)
}

func (animal *Animal) Move(name string) {
	fmt.Printf("The %s %s.\n", name, animal.locomotion)
}

func (animal *Animal) Speak(name string) {
	fmt.Printf("The %s goes '%s'.\n", name, animal.noise)
}

func handleAnimal(animalName string, animals map[string]*Animal) *Animal {
	animal, exists := animals[animalName]
	if !exists {
		fmt.Printf("Unknown animal: %s\n", animalName)
		return nil
	}
	return animal
}

func initAnimals() map[string]*Animal {
	animals := make(map[string]*Animal)
	animals["cow"] = &Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = &Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = &Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	return animals
}
