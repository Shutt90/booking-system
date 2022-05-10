package routes

import (
	controllers "booking-system/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookingsRoutes = func(router *mux.Router) {
	router.HandleFunc("/booking/", controllers.GetBookings)
	// router.HandleFunc("/booking/", controllers.CreateBooking)

}
