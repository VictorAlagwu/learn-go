package main

import (
	"fmt"
)

func main()  {
	half(3)
}

// Reason I used string is because I will learn interfaces in another chapter
func half(x uint) (int, error)  {
	halfValue := x/2
	if halfValue % 2 == 0{
		return fmt.Println(1, true)
	} else {
		return fmt.Println(0, false)
	}
}