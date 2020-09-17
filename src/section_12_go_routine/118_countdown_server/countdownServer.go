package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)
	listenner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		countingDownHandler(conn)
	}

}

func countingDownHandler(conn net.Conn) {
	defer func() {
		io.WriteString(conn, "Your connection will be close by server")
		conn.Close()
	}()
	count := 5
	for {
		io.WriteString(conn, "Hello\n")
		time.Sleep(time.Second)
		count--
		if count == 0 {
			io.WriteString(conn, "Enter number : ")
			// todo
			break
		}
	}

}
