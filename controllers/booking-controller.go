package controllers

import (
	bookings "booking-system/models"
	"encoding/json"
	"net/http"
)

var NewBooking bookings.Booking

func GetBookings (w http.ResponseWriter, r *http.Request) {
	newBooking := bookings.GetAllBookings()
	res, _ := json.Marshal(newBooking)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
