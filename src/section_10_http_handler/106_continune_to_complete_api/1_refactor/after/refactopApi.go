package main

import (
	"log"
	"net/http"
	"os"
)

// type myHandler func(http.ResponseWriter, *http.Request)

// func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	m(w, r)
// }

func inventory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello 2"))
}

func main() {

	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", http.HandlerFunc(inventory))

}
