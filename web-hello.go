package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Running http<h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Running on server :3000")
	http.ListenAndServe(":3000", nil)
}
