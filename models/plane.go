package models

import (
	"time"

	"gorm.io/gorm"
)

type PlaneM struct {
	PlaneId   uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"plane_id"`
	AirlineId uint           `json:"airline_id"`
	PlaneName string         `gorm:"not null" json:"plane_name"`
	PlaneCode string         `gorm:"not null" json:"plane_code"`
	Status    string         `gorm:"not null" json:"status"`
	ImgUrl    string         `gorm:"not null" json:"img_url"`
	Model     string         `gorm:"not null" json:"model"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
