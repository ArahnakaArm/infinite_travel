package models

import (
	"time"

	"gorm.io/gorm"
)

type Airline struct {
	AirlineId     uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"airline_id"`
	AirlineName   string         `gorm:"not null" json:"airline_name" validate:"nonzero"`
	AirlineCode   string         `gorm:"not null" json:"airline_code" validate:"nonzero"`
	PlaneServices []PlaneM       `gorm:"foreignKey:AirlineId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"plane_services"`
	Status        string         `gorm:"not null" json:"status" validate:"regexp=^[a-zA-Z]*$,nonzero"`
	ImgUrl        string         `gorm:"not null" json:"img_url" validate:"nonzero"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
