package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// подключает к серверу на порту 8000
func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	defer fmt.Println("Соединение с сервером закрыто")
	fmt.Println("Соединение с сервером на порту 8000 установлено")

	stdinR := bufio.NewReader(os.Stdin)

	for {
		msg, err := stdinR.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		if string(msg) == "q\n" {
			return
		}

		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Println(err)
			return
		}

		r := bufio.NewReader(conn)

		for {
			msg, _, err := r.ReadLine()
			if err != nil {
				log.Println(err)
				return
			}
			text := string(msg)
			if text == "end" {
				break
			}
			fmt.Println(text)
		}
	}
}
