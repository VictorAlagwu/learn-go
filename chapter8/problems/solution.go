package main

import (
	"fmt"
)

func square (a *float64) {
	*a = *a * *a
	fmt.Println("Address:",a ,"Value:", *a)
}

func swap (xp, yp *int) {
	*xp, *yp = *yp, *xp
}

func main() {
	// How do you get the memory address of a variable => By using & or new
	// How do you assign a value to a pointer => by prefixing it with the * symbol
	// How do you create a new pointer => By using the "new" built-in function
	
	a := 1.5
	square(&a)

	// Swap Numbers
	var x int = 1
	var y int = 2
	fmt.Println("Current values: x=>",x,",y:",y)
	swap(&x, &y)
	fmt.Println("Swap Values: x=>", x , ", y=>", y)
}