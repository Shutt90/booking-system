package Bookings

import (
	"net/http"

	"booking-system/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Booking struct {
	gorm.Model
	Name      string
	Email     string
	Arrival   string
	Departure string
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetAllBookings() []Booking {
	var Booking []Booking
	db.Find(&Booking)
	return Booking
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	db.Create(&Booking{})
}
