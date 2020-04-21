package main

import (
	"fmt"
)

func main() {
	//How do you access the 4th element of an array or slice
	/**
	* Array => arr[3]
	* Slides => sli[3]
	 */
	//Length of a slice => 6
	x := [6]string{"a", "b", "c", "d","e","f"}
	fmt.Println(x[2:5]) // expecting [c,d,e]
}
