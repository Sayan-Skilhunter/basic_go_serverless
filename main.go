package main

import (
	"fmt"
	"io"
	"net/http"
)

func gerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside root handler")
	io.WriteString(w, "This is the root URL")
}

func main() {
    http.HandleFunc("/", gerRoot)
    err := http.ListenAndServe(":8080", nil)

	if err!= nil {
        panic(err)
    } else {
		fmt.Println("Server listening on localhost:8080")
	}
}