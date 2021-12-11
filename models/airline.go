package models

import (
	"time"

	"gorm.io/gorm"
)

type Airline struct {
	AirlineId    string         `gorm:"PRIMARY_KEY" json:"seat_id"`
	AirlineName  string         `gorm:"not null" json:"airline_name"`
	AirlineCode  string         `gorm:"not null" json:"airline_code"`
	PlaneService []PlaneService `gorm:"foreignkey:airline_id" json:"plane_service"`
	Status       string         `gorm:"not null" json:"status"`
	ImgUrl       string         `gorm:"not null" json:"img_url"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
