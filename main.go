package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/yosa12978/gojsonrpc/models"
)

func main() {
	port := 8089
	listen, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1<<10)
	readlen, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		log.Printf("Error with handling request")
		return
	}
	buffer = buffer[:readlen]

	req, batch, err := models.ParseRequest(buffer)
	if err != nil {
		json.NewEncoder(conn).Encode(err)
		return
	}
	if req != nil {
		req.ProcRequest(Playground{})
	} else {
		batch.ProcBatch(Playground{})
	}

}

func ParseRequest() {

}

type Playground struct {
}

func (pg *Playground) Add(a int, b int) int {
	return a + b
}
