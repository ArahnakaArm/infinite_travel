package models

import (
	"time"

	"gorm.io/gorm"
)

type PlaneM struct {
	PlaneId   uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"plane_id"`
	AirlineId uint           `json:"airline_id" validate:"nonzero"`
	Airline   Airline        `json:"airline" validate:"nonzero"`
	PlaneName string         `gorm:"not null" json:"plane_name" validate:"nonzero"`
	PlaneCode string         `gorm:"not null" json:"plane_code" validate:"nonzero"`
	Status    string         `gorm:"not null" json:"status" validate:"regexp=^[a-zA-Z]*$,nonzero"`
	ImgUrl    string         `gorm:"not null" json:"img_url" validate:"nonzero"`
	Model     string         `gorm:"not null" json:"model" validate:"nonzero"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
