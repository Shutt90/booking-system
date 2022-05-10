package routes

import (
	controllers "booking-system/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookingsRoutes = func(router *mux.Router) {
	router.HandleFunc("/booking/", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/booking/", controllers.CreateBooking).Methods("POST")

}
