package models

type Flight struct {
	FlightId   string `gorm:"PRIMARY_KEY" json:"flight_id"`
	FlightName string ` json:"flight_name"`
}
