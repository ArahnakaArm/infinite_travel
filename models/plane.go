package models

type PlaneM struct {
	PlaneId   string `gorm:"PRIMARY_KEY" json:"plane_id"`
	AirlineId string `gorm:"column:airline_id;not null" json:"airline_id"`
	PlaneName string `gorm:"not null" json:"plane_name"`
	PlaneCode string `gorm:"not null" json:"plane_code"`
	Status    string `gorm:"not null" json:"status"`
	ImgUrl    string `gorm:"not null" json:"img_url"`
	Model     string `gorm:"not null" json:"model"`
}
