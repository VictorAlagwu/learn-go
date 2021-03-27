package main

import (
	"log"
	"os"

	_ "github.com/victoralagwu/learn-go/goinaction/chapter2/matchers"
	"github.com/victoralagwu/learn-go/goinaction/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
