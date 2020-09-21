package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello Docker world!")
}
