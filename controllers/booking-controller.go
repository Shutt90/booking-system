package controllers

import (
	bookings "booking-system/models"
	"booking-system/utils"
	"encoding/json"
	"net/http"
)

var NewBooking bookings.Booking

func GetBookings(w http.ResponseWriter, r *http.Request) {
	newBooking := bookings.GetAllBookings()
	res, _ := json.Marshal(newBooking)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	CreateBooking := &bookings.Booking{}
	println(CreateBooking)
	utils.ParseBody(r, CreateBooking)
	b := CreateBooking.CreateBooking()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
