package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", GetHello)

	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server closed\n")
		} else {
			fmt.Printf("Error starting server: %s\n", err)
		}
		os.Exit(1)
	}
	fmt.Print("Server is running...")
}
