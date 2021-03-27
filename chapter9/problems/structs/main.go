package main

import (
	c "github.com/victoralagwu/learn-go/chapter9/problems/structs/computer"
	"fmt"
)

func main() {
	spec := c.Spec {
		Maker: "apple",
		Price: 50000,
	}
	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}