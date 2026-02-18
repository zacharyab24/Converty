package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	// Serve files from the ./static directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Printf("Server listening on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
