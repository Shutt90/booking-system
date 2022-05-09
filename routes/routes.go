package routes

import (
	"github.com/gorilla/mux"
	"booking-system/controllers"
)

var RegisterBookingsRoutes  = func(router *mux.Router) {
	router.HandleFunc("/booking/", controllers.CreateBooking)
	router.HandleFunc("/booking/", controllers.GetBookings)

}