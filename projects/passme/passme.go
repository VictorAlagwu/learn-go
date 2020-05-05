package main

import (
	"os"
	"fmt"
)
const (
	errPwd = "Sorry, wrong password entered for user %q.\n"
	accessDenied = "Access denied for user %q.\n"
	accessGiven = "Hello, %s Welcome to PASSME\n"
)
func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	u,p := args[1], args[2]
	if u != "victor" {
		fmt.Printf(accessDenied, u)
	}else if p != "password" {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessGiven, u)
	}
}