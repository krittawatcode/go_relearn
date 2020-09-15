package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func inventory(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/items":
		w.Write([]byte("handle items"))
	case "/price":
		w.Write([]byte("handle price"))
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "sorry, no such page: %s", r.URL)
	}
}

func main() {

	log.SetFlags(0)
	log.Println(os.Getpid())
	http.ListenAndServe(":8080", http.HandlerFunc(inventory))

}
