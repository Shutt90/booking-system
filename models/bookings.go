package Bookings

import (
	"booking-system/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Booking struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	Arrival   string `json:"arrival"`
	Departure string `json:"departure"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Booking{})

}

func GetAllBookings() []Booking {
	var Booking []Booking
	db.Find(&Booking)
	return Booking
}

func (b *Booking) CreateBooking() *Booking {
	db.Create(&b)
	return b
}
