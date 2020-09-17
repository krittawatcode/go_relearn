package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Ltime)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Connected to server successfully")
	log.Println("Copy from conn to stdout")
	// copy conn --> stdout
	go io.Copy(os.Stdout, conn)
	log.Println("Copy from stdout to conn")
	io.Copy(conn, os.Stdin)
}
