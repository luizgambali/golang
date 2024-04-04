package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorld)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Olá Mundo, de um container Docker rodando Go!")
}
