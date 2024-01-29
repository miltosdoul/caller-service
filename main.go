package main

import (
	"caller/pkg/client"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var counterClient *client.CounterClient = client.NewCounterClient()

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	reqCount := counterClient.GetRequestCount()
	io.WriteString(w, reqCount)
}
