package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {

	fs := http.FileServer(http.Dir("./frontend/"))
	http.Handle("/", fs)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
