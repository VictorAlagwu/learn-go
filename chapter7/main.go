package main

import (
	"fmt"
)



func main() {
	firstCall()

	
	x, y := f()
	fmt.Println(x , y)
	fmt.Println(add(1,2,3))
	answer := average([]float64{10, 10, 10, 10, 10})
	fmt.Println("The answer is",answer);
	
	// Defer
	defer lastCall()

	//Closure
	createEven := evenNumGenerator()
	fmt.Println("Closure: ", createEven())
	fmt.Println("Closure:", createEven())

	//Factorial
	fmt.Println("Factorial:", factorial(0))
}


func average(x []float64) float64 {
	total := 0.0
	for _, v:= range x {
		total += v // v == x[i]
	}
	return total / float64(len(x))
}
func f() (int, string) {
	return 32, "Victor"
}

// Using Spread for the params => In this case, we take zero or more ints
func add(params ...int) int {
	total := 0
	for _,v := range params {
		total += v
	}
	return total
}
//Recursion: Takes in unsigned integer, and returns unsigned integer
func factorial(x uint) uint  {
	if x==0 {
		return 1
	}	
	return x * factorial(x-1)
}

// Closure Usage
func evenNumGenerator() func () uint  {
	x := uint(0)
	return func () (ret uint) {
		ret = x
		x += 2
		return
	}
}

func firstCall () {
	fmt.Println("This is the first Function Call")
}

func lastCall () {
	defer func () {
		str := recover()
		fmt.Println(str)
		fmt.Println("This is the last Function call")
	}()
	panic("Panic mood")
}

