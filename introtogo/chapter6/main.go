package main

import (
	"fmt"
)

func main() {
	//Array
	x := [5]float64{10, 10, 10, 10, 10}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println("Array")
	fmt.Println(total / float64(len(x)))
	//Slides
	arr := []string{"Victor", "Tester", "Sample", "Baby"}
	y := append(arr, "Chim", "Amaka")
	fmt.Println("Slides")
	fmt.Println(arr, y)
	//Maps
	newMap := make(map[string]int)
	newMap["first"] = 20
	newMap["second"] = 12
	delete(newMap, "first")
	fmt.Println("***Map Vibe***")
	sam, ok := newMap["first"]
	fmt.Println(sam, ok)
	fmt.Println(newMap["first"])
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
