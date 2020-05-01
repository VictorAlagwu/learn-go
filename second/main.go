package main

import (
	"strings"
	// "strconv"
	"os"
	"fmt"
	"reflect"
)
type Newtype string

func main () {
    args := os.Args[1]
    var name Newtype = "Hsd"     	 
	l := len(args)
	value := strings.Repeat("!", l)
	s := value + args + value 
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(name))
	// feet, _ := strconv.ParseFloat(args, 64)

	// meters := feet * 0.3048
	// fmt.Printf("%f feet is %f meters.\n", feet,meters)
}