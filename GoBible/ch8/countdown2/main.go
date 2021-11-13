package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan int)
	complete := make(chan int)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- 1
	}()

	go func() {
		tick := time.Tick(1 * time.Second)
		for n := 10; n > 0; n-- {
			fmt.Println(n)
			<-tick
		}
		complete <- 1
	}()

	fmt.Println("commencing countdown.")
	select {
	case <-complete:
	case <-abort:
		fmt.Println("launch abort!")
		return
	}
	launch()
}

func launch() {
	fmt.Println("LAUNCH!")
}
