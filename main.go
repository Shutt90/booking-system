package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleDate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parser() err: %v", err)
	}

}

func main() {
	fs := http.FileServer(http.Dir("./src/assets"))
	http.Handle("/", fs)
	http.HandleFunc("/book", handleDate)
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Starting server at port 8080/n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
