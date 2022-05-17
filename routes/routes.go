package routes

import (
	controllers "booking-system/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterBookingsRoutes = func(router *mux.Router) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
	router.HandleFunc("/bookings/", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBooking).Methods("POST")

}
