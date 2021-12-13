package models

import (
	"time"

	"gorm.io/gorm"
)

type Airport struct {
	AirportId   uint   `gorm:"PRIMARY_KEY;autoIncrement:false" json:"airport_id"`
	AirportName string `gorm:"not null" json:"airport_name"`
	AirportCode string `gorm:"not null" json:"airport_code"`
	/* 	AirlineServices []Airline      `gorm:"many2many:airport_airline;" json:"airline_services"` */
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
