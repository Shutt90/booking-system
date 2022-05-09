package routes

import (
	"github.com/gorilla/mux"
	controllers "booking-system/controllers"
)

var RegisterBookingsRoutes  = func(router *mux.Router) {
	router.HandleFunc("/booking/", controllers.GetBookings)
	router.HandleFunc("/booking/", controllers.CreateBooking)


}