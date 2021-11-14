package models

type Explore struct {
	ExploreId uint             `gorm:"PRIMARY_KEY"`
	Header    string           `gorm:"not null" json:"header"`
	Content   []ExploreContent `gorm:"foreignkey:explore_id" json:"content"`
}

type ExploreContent struct {
	ExploreContentId uint   `gorm:"PRIMARY_KEY"`
	ExploreId        uint   `gorm:"column:explore_id"`
	Header           string `gorm:"not null" json:"header"`
	Body             string `gorm:"not null" json:"body"`
	ImageUrl         string `gorm:"not null" json:"image_url"`
}

type ExploreRequest struct {
	Header string `gorm:"not null" json:"header"`
}

type ExploreContentRequest struct {
	ExploreId uint   `gorm:"column:explore_id"`
	Header    string `gorm:"not null" json:"header"`
	Body      string `gorm:"not null" json:"body"`
	ImageUrl  string `gorm:"not null" json:"image_url"`
}
