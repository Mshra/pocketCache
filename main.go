package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handleInsert(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website")
}

func handleDelete(w http.ResponseWriter, r *http.Request) {}

func handleSearch(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/insert", handleInsert)
	http.HandleFunc("/delete", handleInsert)
	http.HandleFunc("/search", handleInsert)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
