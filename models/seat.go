package models

import (
	"time"

	"gorm.io/gorm"
)

type Seat struct {
	SeatId       uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"seat_id"`
	FlightId     uint           `json:"flight_id"`
	SeatNumber   string         `gorm:"not null;unique" json:"seat_number"`
	TicketNumber string         `gorm:"default:null;unique" json:"ticket_number"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
