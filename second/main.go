package main

import (
	"strconv"
	"os"
	"fmt"
)

func main () {
    args := os.Args[1]
	 
	feet, _ := strconv.ParseFloat(args, 64)

	meters := feet * 0.3048
	fmt.Printf("%f feet is %f meters.\n", feet,meters)
}