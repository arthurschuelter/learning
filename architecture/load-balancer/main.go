package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	for i := 1; i <= 3; i++ {
		serverNumber := i
		port := 8080 + i
		go func() {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "<h1>Server %d</h1>", serverNumber)
			})
			address := fmt.Sprintf(":%d", port)
			log.Printf("Server %d listening on %s", serverNumber, address)
			log.Fatal(http.ListenAndServe(address, handler))
		}()
	}

	select {}
}
