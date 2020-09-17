package main

import "net"

func main() {
	net.Dial("tcp", "localhost:8080")
}
