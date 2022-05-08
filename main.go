package main

import (
	Bookings "booking-system/models"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {

	fs := http.FileServer(http.Dir("./frontend/"))
	http.Handle("/", fs)
	http.HandleFunc("/book", Bookings.CreateBooking).Methods("POST")
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
