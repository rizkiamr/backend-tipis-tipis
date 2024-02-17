package main

import (
	"fmt"
	"log"
	"net/http"
)

func queryName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, "Kasih paham %s", name)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tipis-Tipis Captt"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/v1/kasih-paham", queryName)

	log.Println("Starting web server at :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
