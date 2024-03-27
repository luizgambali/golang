package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/api/v1/hello", http.HandlerFunc(hello))

	fmt.Println("Server is running at :3000")
	http.ListenAndServe(":3000", mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
