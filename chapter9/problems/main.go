package main

import (
	"fmt"
)

type Employer struct {
	companyName string
}
//Employee :
type Employee struct {
	firstName, lastName string
	age int
	salary int
	Employer
}
func main() {
	emp1 := &Employee{
		firstName: "Victor", 
		lastName: "Alagwu", 
		Employer: Employer {
			"TM30 Global Limited",
		},
	}
	// emp2 := &Employee{
	// 	"Vic", "Ala", 6,4, Employer: Employer {
	// 		"TM30 Global Limited",
	// 	},
	// }
	fmt.Println("Employee: ", emp1.companyName)	
}