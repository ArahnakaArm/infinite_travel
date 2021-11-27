package models

import (
	"time"

	"gorm.io/gorm"
)

type Explore struct {
	ExploreId uint             `gorm:"PRIMARY_KEY" json:"exploreId"`
	Title     string           `gorm:"not null" json:"title"`
	Author    string           `gorm:"not null" json:"author"`
	Paragraph string           `gorm:"not null" json:"paragraph"`
	ImageUrl  string           `gorm:"not null" json:"image_url"`
	Content   []ExploreContent `gorm:"foreignkey:explore_id" json:"content"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
}

type ExploreContent struct {
	ExploreContentId uint           `gorm:"PRIMARY_KEY"`
	ExploreId        uint           `gorm:"column:explore_id"`
	Title            string         `gorm:"not null" json:"title"`
	Body             string         `gorm:"not null" json:"body"`
	ImageUrl         string         `gorm:"not null" json:"image_url"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ExploreRequest struct {
	Title     string `gorm:"not null" json:"title"`
	Author    string `gorm:"not null" json:"author"`
	Paragraph string `gorm:"not null" json:"paragraph"`
	ImageUrl  string `gorm:"not null" json:"image_url"`
}

type ExploreContentRequest struct {
	ExploreId uint   `gorm:"column:explore_id"`
	Title     string `gorm:"not null" json:"title"`
	Body      string `gorm:"not null" json:"body"`
	ImageUrl  string `gorm:"not null" json:"image_url"`
}
