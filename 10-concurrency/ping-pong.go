package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	begin, ping, pong, stop = "begin", "ping", "pong", "stop"
)

var game chan string
var m sync.Mutex
var count = 0
var chance int

func main() {
	game = make(chan string)
	wg := sync.WaitGroup{}

	// 1 из 5 раз кому-то должно повезти
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance = r.Intn(5) + 1

	wg.Add(1)
	go func() {
		defer wg.Done()

		for val := range game {
			hit("p1", val)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for val := range game {
			hit("p2", val)
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
		fmt.Println(p, ": ", ping)
		if checkChance() {
			close(game)
			fmt.Println("winner: ", p)
			return
		}
		game <- ping
	case ping:
		fmt.Println(p, ": ", pong)
		if checkChance() {
			close(game)
			fmt.Println("winner: ", p)
			return
		}
		game <- pong
	}
}

func checkChance() bool {
	gameOver := false
	m.Lock()
	count += 1
	if count == chance {
		gameOver = true
	}
	m.Unlock()
	return gameOver
}
