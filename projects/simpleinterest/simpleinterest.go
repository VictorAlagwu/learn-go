package simpleinterest

import "fmt"

func init() {
	fmt.Println("Simple Interest Package Initialized")
}

//Calculate : Get interest 
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}