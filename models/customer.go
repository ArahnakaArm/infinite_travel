package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerId   uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"customer_id"`
	UserName     string         `gorm:"not null" json:"user_name" validate:"nonzero"`
	Password     string         `gorm:"not null" json:"password" validate:"nonzero"`
	FirstName    string         `gorm:"not null" json:"first_name" validate:"nonzero"`
	LastName     string         `gorm:"not null" json:"last_name" validate:"nonzero"`
	MiddleName   string         `json:"middle_name"`
	IdCard       string         `gorm:"not null;unique" json:"id_card"  validate:"nonzero,min=13,max=13,regexp=^[0-9]*$"`
	VisaNumber   string         `gorm:"unique" json:"visa_number" validate:"nonzero"`
	MobileNumber string         `gorm:"not null;unique" json:"mobile_number" validate:"nonzero,min=11,max=11,regexp=^[0-9]*$"`
	Nation       string         `gorm:"not null" json:"nation" validate:"nonzero"`
	Gender       string         `gorm:"not null" json:"gender" validate:"nonzero"`
	Tickets      []Ticket       `gorm:"foreignKey:CustomerId" json:"ticket"`
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
