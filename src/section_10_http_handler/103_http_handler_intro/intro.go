package main

import "net/http"

func main() {
	/* --- way 1 --- */
	// http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// }))

	/* --- way 2 --- */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello 2"))
	})
	http.ListenAndServe(":8080", nil)
}
