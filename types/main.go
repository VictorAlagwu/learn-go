package main

import (
	"time"
	"fmt"
)

type gram float64
type ounce float64

func main() {
	var g gram = 1000
	var o ounce 
	o = ounce(g) * 0.0335274

	fmt.Printf("%g grams is %.2f ounces\n", g, o)

	t := time.Second * 10
	fmt.Println(t)
	i := 10
	t = time.Second * time.Duration(i)
	fmt.Println(t) 
	const c = 10
	t = time.Second * c
	fmt.Println(t) 
}