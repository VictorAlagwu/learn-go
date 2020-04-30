package main

import (
	"fmt"
)

// FahrenriteToCelsius: 
func fahrenriteToCelsius (x float64) float64 {
	var output float64 = (x - 32) * 5 / 9
	return output
}
 
// CelsiusToFahrenrite
func celsiusToFahrenrite (x float64) float64 {
	var output float64 = (x * 9/5) + 32
	return output
}

func feetToMeter (x float64) float64 {
	var output float64 = x/3.281
	return output
}

func arrayChecker(a int, arr [3]int) bool{
	for _, b := range arr {
		if b == a {
			return true
		} 
	} 	
	return false
}

func main() {
	var selected  int;
	var value float64;
	acceptedSelectors := [3]int{1, 2, 3}
	fmt.Println("Hello, Welcome to the Terminal Converter")
	fmt.Println(`
		Select your conversion option:
		1: Fahrenrite to Celsius,
		2: Celesius to Fahrenrite,
		3: Feet to Meter
		`)
	fmt.Scanf("%d", &selected)
	status := arrayChecker(selected, acceptedSelectors)
	if status != true {
		 fmt.Println("Invalid value selected", status)
	}
	fmt.Println("Please enter the value, you will like to convert")
	
	fmt.Scanf("%f", &value)
	if selected == 1 {
		result := fahrenriteToCelsius(value)
		fmt.Println(value, " Fahrenrite in Celsius is", result)
	} else if selected == 2 {
		result := celsiusToFahrenrite(value)
		fmt.Println(value, " Celsius in Fahrenrite is", result)
	} else if selected == 3 {
		result := feetToMeter(value) 
		fmt.Println(value, " Feet in Meter is", result)
	} 

}
