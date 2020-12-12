package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	s := http.Server{
		Addr:    ":3000",
		Handler: http.DefaultServeMux,
	}
	s.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
