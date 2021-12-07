package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerId   string         `gorm:"PRIMARY_KEY" json:"customer_id"`
	FirstName    string         `gorm:"not null" json:"first_name"`
	LastName     string         `gorm:"not null" json:"last_name"`
	MiddleName   string         `json:"middle_name"`
	IdCard       string         `gorm:"not null;unique" json:"id_card"`
	VisaNumber   string         `gorm:"unique" json:"visa_number"`
	MobileNumber string         `gorm:"not null;unique" json:"mobile_number"`
	Nation       string         `gorm:"not null" json:"nation"`
	Gender       string         `gorm:"not null" json:"gender"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	/* 	Title     string           `gorm:"not null" json:"title"`
	   	Author    string           `gorm:"not null" json:"author"`
	   	Paragraph string           `gorm:"not null" json:"paragraph"`
	   	ImageUrl  string           `gorm:"not null" json:"image_url"`
	   	Content   []ExploreContent `gorm:"foreignkey:explore_id" json:"content"`
	   	CreatedAt time.Time        `json:"created_at"`
	   	UpdatedAt time.Time        `json:"updated_at"`
	   	DeletedAt gorm.DeletedAt   `gorm:"index" json:"deleted_at"` */
}

/* type ExploreContent struct {
	ExploreContentId uint           `gorm:"PRIMARY_KEY"`
	ExploreId        uint           `gorm:"column:explore_id"`
	Title            string         `gorm:"not null" json:"title"`
	Paragraph        string         `gorm:"not null" json:"paragraph"`
	ImageUrl         string         `gorm:"not null" json:"image_url"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
*/
