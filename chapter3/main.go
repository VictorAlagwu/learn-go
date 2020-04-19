package main

import (
	"fmt"
)

func main() {
	var (
		// name            string
		// age             int
		favouriteNumber float64
	)
	fmt.Println("Hello, Welcome to the World Cinema")
	// fmt.Println("Please, enter your name?: ")
	// fmt.Scanf("%f", &name)
	// fmt.Println("What is your age?:")
	// fmt.Scanf("%f", &age)
	fmt.Println("Enter your favourite Number?:")
	fmt.Scanf("%f", &favouriteNumber)
	var output float64 = (favouriteNumber - 32) * 5 / 9
	fmt.Println("Favouite number", output)
}
