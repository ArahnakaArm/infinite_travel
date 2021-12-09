package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	TicketId     string         `gorm:"PRIMARY_KEY" json:"ticket_id"`
	CustomerId   string         `gorm:"column:customer_id;not null" json:"customer_id"`
	TicketNumber string         `gorm:"not null" json:"ticket_number"`
	FlightId     string         `gorm:"column:flight_id" json:"flight_id"`
	Flight       Flight         `json:"flight"`
	AirlineId    string         `gorm:"column:airline_id" json:"airline_id"`
	PlaneId      string         `gorm:"column:plane_id" json:"plane_id"`
	Seat         string         `gorm:"not null;unique" json:"seat"`
	Status       string         `gorm:"not null" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
