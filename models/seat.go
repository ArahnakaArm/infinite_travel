package models

import (
	"time"

	"gorm.io/gorm"
)

type Seat struct {
	SeatId     string         `gorm:"PRIMARY_KEY" json:"seat_id"`
	FlightId   string         `gorm:"column:flight_id;not null" json:"flight_id"`
	SeatNumber string         `gorm:"not null" json:"seat_number"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
