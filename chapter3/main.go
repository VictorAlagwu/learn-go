package main

import (
	"math"
	"fmt"
)

// type Converter float64
//Using defined types as results
const (
	feetInMeters float64 = 3.281
    feetInYards  = feetInMeters / 0.9144 //Follows the type of the last constant float64 -feetInMeters
)
// FahrenriteToCelsius: 
func fahrenriteToCelsius (x float64) float64 {
	var output float64 = (x - 32) * 5 / 9
	return output
}
 
// CelsiusToFahrenrite
func celsiusToFahrenrite (x float64) float64 {
	var o float64 = (x * 9/5) + 32
	return o
}

func feetToMeter (x float64) float64 {
	var o float64 = math.Round(x * feetInMeters)
	return o
}

func feetToYard(x float64) float64 {
	var o float64 = math.Round(x * feetInYards)
	return o
}

func arrayChecker(a int, arr [4]int) bool{
	for _, b := range arr {
		if b == a {
			return true
		} 
	} 	
	return false
}

func main() {
	var s  int;
	var v float64;
	xs := [4]int{1, 2, 3, 4}
	fmt.Println("Hello, Welcome to the Terminal Converter")
	fmt.Println(`
		Select your conversion option:	
		1: Fahrenrite to Celsius,
		2: Celesius to Fahrenrite,
		3: Feet to Meter,
		4: Feet In Yard
		`)
	fmt.Scanf("%d", &s)
	status := arrayChecker(s, xs)
	if status != true {
		 fmt.Println("Invalid value selected", status)
	}
	fmt.Println("Please enter the value, you will like to convert")
	
	fmt.Scanf("%f", &v)
	if s == 1 {
		r := fahrenriteToCelsius(v)
		fmt.Printf("%f fahrenrite is %f celsius.\n",v,r)
	} else if s == 2 {
		r := celsiusToFahrenrite(v)
		fmt.Printf("%f celsius is %f fahrenrite.\n",v,r)
	} else if s == 3 {
		r := feetToMeter(v) 
		fmt.Printf("%f feet is %f meters.\n",v,r)
	} else if s == 4 {
		r := feetToYard(v) 
		fmt.Printf("%f feet is %f yards.\n",v,r)
	}

}
