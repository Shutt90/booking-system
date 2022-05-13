package main

import (
	"booking-system/config"
	"booking-system/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	fs := http.FileServer(http.Dir("./frontend/"))

	r := mux.NewRouter()
	routes.RegisterBookingsRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe("localhost:8080", fs))

	db = config.GetDB()

}
