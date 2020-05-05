package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if n, err := strconv.Atoi(os.Args[1]); err == nil {
		fmt.Println("Converted number :", n)
		fmt.Println("Returned error value :", err)
	} else {
		fmt.Printf("error: %q input has an error.\n", n)
	}

	
}