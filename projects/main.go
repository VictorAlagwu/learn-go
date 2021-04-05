package main

import (
	"fmt"
	"log"

	"github.com/victoralagwu/learn-go/projects/simpleinterest"
	_ "github.com/victoralagwu/learn-go/projects/simpleinterest"
)

var p, r , t = 5000.0, 10.0, 1.0

func init() {
	println("Main Package Initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	} 
	if r < 0 {
		log.Fatal("Rate is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}
func main() {
	fmt.Println("::>>Simple Interest Calculator<<::")

	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple Interest is", si)
}