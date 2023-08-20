package netsrv

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sort"

	"go-core-4/11-net/search-engine/pkg/index"
)

func Start(documents []index.Document, index index.Index) {
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
		go handler(conn, documents, index)
	}

}

// обработчик подключения
func handler(conn net.Conn, documents []index.Document, index index.Index) {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		res := searchString(string(msg), documents, index)

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

// поиск строки из запроса в отсканированных документах
func searchString(s string, documents []index.Document, index index.Index) (result []string) {
	for _, id := range index[s] {
		i := sort.Search(len(documents), func(i int) bool {
			return documents[i].ID >= id
		})
		if i < len(documents) && documents[i].ID == id {
			result = append(result, documents[i].Title)
		}
	}
	return result
}
