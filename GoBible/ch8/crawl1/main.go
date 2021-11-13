package main

import (
	"Learning-Golang/GoBible/ch5/links"
	"fmt"
	"log"
	"os"
)

var tokens = make(chan int, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- 1
	list, err := links.Extract(url)
	<-tokens

	if err != nil {
		log.Printf("crawl failed:%s", err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int

	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}

}
