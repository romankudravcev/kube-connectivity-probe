package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	targetHost := os.Getenv("TARGET_HOST")

	if port == "" {
		fmt.Println("PORT must be set")
		return
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	go http.ListenAndServe(":"+port, nil)
	fmt.Printf("Server listening on port %s\n", port)

	if targetHost != "" {
		for {
			address := fmt.Sprintf("%s:%s", targetHost, port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("Failed to connect to %s - %v\n", address, err)
			} else {
				fmt.Printf("Successfully connected to %s\n", address)
				conn.Close()
			}
			time.Sleep(10 * time.Second)
		}
	}
}
