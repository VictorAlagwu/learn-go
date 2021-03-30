package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){ 
	for _, url := range os.Args[1:] {
		var newUrl string
		if strings.HasPrefix(url, "http://") {
			newUrl = url
		} else if strings.HasPrefix(url, "https://") {
			newUrl = url
		} else {
			newUrl = "http://" + url
		}
		fmt.Printf("%s", newUrl  + " " + url)
		resp, err := http.Get(newUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", b)
	}
}