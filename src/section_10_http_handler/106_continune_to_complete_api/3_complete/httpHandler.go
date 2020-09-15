package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type inventory map[string]float64

func (iv inventory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/items":
		for k, v := range iv {
			fmt.Fprintf(w, "%s : %.2f\n", k, v)
		}
	case "/price":
		searchItem := r.URL.Query().Get("items")
		price, ok := iv[searchItem]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no item: %s", searchItem)
			return
		}
		fmt.Fprintf(w, "%.2f\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "sorry, no such page: %s", r.URL)
	}
}

func main() {

	log.SetFlags(0)
	log.Println(os.Getpid())
	inven := inventory{
		"apple":  1.25,
		"orange": 0.99,
	}
	http.ListenAndServe(":8080", inven)

}
