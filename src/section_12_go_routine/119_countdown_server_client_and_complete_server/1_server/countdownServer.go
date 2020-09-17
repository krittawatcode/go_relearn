package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
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
		log.Println("New client is connect : ", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
			return
		}
		go countingDownHandler(conn)
	}

}

func countingDownHandler(conn net.Conn) {
	defer func() {
		io.WriteString(conn, "Your connection will be close by server") // write back to the one who connect
		log.Println("Connection is closed be client")
		conn.Close()
	}()
	io.WriteString(conn, "Enter number : ")
	input := bufio.NewScanner(conn)
	count := Scan(input)
	if count > 20 {
		return
	}
	for {
		io.WriteString(conn, strconv.Itoa(count)+"\n") // write back to the one who connect
		time.Sleep(time.Second)
		count--
		if count < 0 {
			// write back to the one who connect
			// todo
			io.WriteString(conn, "Enter number : ")
			count = Scan(input)
			if count == 0 {
				break
			}
		}
	}

}

func Scan(input *bufio.Scanner) int {
	if ok := input.Scan(); !ok {
		log.Println("Cannot scan value from conn")
		log.Println("Connection is closed be client")
		return 0
	}
	count, err := strconv.Atoi(input.Text())
	if err != nil {
		log.Println("Cannot convert value from Text to int. Value is : ", input.Text())
		return 0
	}
	return count
}
