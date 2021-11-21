package models

import (
	"time"

	"gorm.io/gorm"
)

type Explore struct {
	ExploreId uint             `gorm:"PRIMARY_KEY" json:"exploreId"`
	Header    string           `gorm:"not null" json:"header"`
	Author    string           `gorm:"not null" json:"author"`
	Content   []ExploreContent `gorm:"foreignkey:explore_id" json:"content"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
}

type ExploreContent struct {
	ExploreContentId uint           `gorm:"PRIMARY_KEY"`
	ExploreId        uint           `gorm:"column:explore_id"`
	Header           string         `gorm:"not null" json:"header"`
	Body             string         `gorm:"not null" json:"body"`
	ImageUrl         string         `gorm:"not null" json:"image_url"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ExploreRequest struct {
	Header string `gorm:"not null" json:"header"`
	Author string `gorm:"not null" json:"author"`
}

type ExploreContentRequest struct {
	ExploreId uint   `gorm:"column:explore_id"`
	Header    string `gorm:"not null" json:"header"`
	Body      string `gorm:"not null" json:"body"`
	ImageUrl  string `gorm:"not null" json:"image_url"`
}
