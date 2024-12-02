package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 3000, "Port to listen and serve.")

	flag.Parse()

	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))
}
