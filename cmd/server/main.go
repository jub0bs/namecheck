package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /hello", handleHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, world!")
}
