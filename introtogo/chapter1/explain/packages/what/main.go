package main

import (
	"fmt"
	"runtime"
)

const sample = "outside scope"

func main() {
	// const sample = "Main scope"
	// var name = "Victor"
	fmt.Println(runtime.NumCPU())
	// hey()
}
