package main

import (
	"Learning-Golang/GoBible/ch9/memo1/memo"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := memo.New(httpGetBody)

	incomingURLs := make(chan string)

	go func() {
		for _, url := range os.Args[1:] {
			incomingURLs <- url
		}
		close(incomingURLs)
	}()

	Start := time.Now()

	for url := range incomingURLs {
		value, err := m.Get(url)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s, %d bytes\n", url, len(value.([]byte)))
	}

	fmt.Printf("completed in %s.\n", time.Since(Start))
}
