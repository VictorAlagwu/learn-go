package main

import (
	"fmt"
)

func main()  {
	x := uint(2)
	fmt.Println(half(x))
}

// Reason I used string is because I will learn interfaces in another chapter
func half(x uint) (uint, bool)  {
	halfValue := x/2
	if x % 2 == 0{
		return halfValue, true
	} else {
		return halfValue, false
	}
}