package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	FlightId   uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"flight_id"`
	FlightName string         `gorm:"not null" json:"flight_name"`
	PlaneMId   uint           `json:"plane_id"`
	PlaneM     PlaneM         `json:"plane_serve"`
	AirlineId  uint           `json:"airline_id"`
	Airline    Airline        `json:"airlined"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
