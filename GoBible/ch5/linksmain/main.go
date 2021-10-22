package main

import (
	"Learning-Golang/GoBible/ch5/links"
	"fmt"
	"log"
	"os"
)

func main() {
	// urls := os.Args[1:]
	urls := []string{"https://www.vilipix.com/"}
	for _, url := range urls {
		links, err := links.Extract(url)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		for i, link := range links {
			fmt.Printf("%3d   %s\n", i, link)
		}
	}
}
