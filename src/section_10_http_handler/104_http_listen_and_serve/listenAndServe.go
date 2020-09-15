package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe("", nil) // will auto redirect to 8080

	// netstat -ano | grep 3752
	// http.ListenAndServe(":8082", nil)
	// http.ListenAndServe(":", nil) // will random port
}
