package main

import (
	"fmt"
	"net/http"
)

type Name string
type Value string
type VValue string
type VVValue string
type VVVValue string
type XXX string

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", "World")
}
