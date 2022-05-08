package Bookings

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"booking-system/config"
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
	db.AutoMigrate(&Booking{})
}

func Bookings() {
	db, err := gorm.Open("mysql", "root:@/booking?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Issue with db connect")
	}
	defer db.Close()
	db.AutoMigrate(&Booking{})

}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	db.Create(&Booking{})

}
