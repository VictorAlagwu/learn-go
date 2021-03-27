package main

import (
	"math/rand"
	"time"
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan<- string) {
	for i := 0; ;i++ {
		c <- "ping"
	}
}

func printer(c <-chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func ponger(c chan<- string){
	for i:=0; ; i++ {
		c <- "pong"
	}
}
func main() {
	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// }

	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	go ponger(c)

	var input string
	fmt.Scanln(&input)
}