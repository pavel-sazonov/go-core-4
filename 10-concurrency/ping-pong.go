package main

import (
	"fmt"
	"sync"
)

const (
	begin, ping, pong, stop = "begin", "ping", "pong", "stop"
)

var game chan string

// var m sync.Mutex
// var msg string

func main() {
	game = make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for val := range game {
			hit("p1: ", val)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for val := range game {
			hit("p2: ", val)
		}
	}()

	fmt.Println("main: ", begin)
	game <- begin

	wg.Wait()
	fmt.Println("main: ", stop)
}

func hit(p string, v string) {
	switch v {
	case begin, pong:
		fmt.Println(p, ping)
		game <- ping
	case ping:
		fmt.Println(p, pong)
		game <- pong
	}
}
