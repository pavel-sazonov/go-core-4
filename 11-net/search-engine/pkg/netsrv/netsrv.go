package netsrv

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"go-core-4/11-net/search-engine/pkg/index"
)

func Start(documents []index.Document) {
	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Сервер слушает на порту: 8000")

	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("клиент подключился")
		go handler(conn, documents)
	}

}

// обработчик подключения
func handler(conn net.Conn, documents []index.Document) {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		res := index.Search(string(msg), documents)

		if len(res) == 0 {
			_, err = conn.Write([]byte("ничего не найдено\n"))
			if err != nil {
				log.Println(err)
				return
			}
		}

		for _, s := range res {
			data := []byte(s)
			data = append(data, '\n')
			_, err := conn.Write(data)
			if err != nil {
				log.Println(err)
				return
			}
		}

		_, err = conn.Write([]byte("end\n"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
