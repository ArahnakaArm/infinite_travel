package models

import (
	"time"

	"gorm.io/gorm"
)

/* type Datetime struct {
	time.Time
}

func (t *Datetime) UnmarshalJSON(input []byte) error {
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006-01-02", strInput)
	if err != nil {
		return err
	}

	t.Time = newTime
	return nil
} */

type Flight struct {
	FlightId             uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"flight_id"`
	FlightName           string         `gorm:"not null" json:"flight_name"`
	PlaneMId             uint           `json:"plane_id"`
	PlaneM               PlaneM         `json:"plane_serve"`
	AirlineId            uint           `json:"airline_id"`
	Airline              Airline        `json:"airline"`
	DepartTime           string         `gorm:"not null" json:"depart_time"`
	ArriveTime           string         `gorm:"not null" json:"arrive_time"`
	DestinationAirportId uint           `json:"destination_airport_id"`
	DestinationAirport   Airport        `json:"destination_airport"`
	OriginAirportId      uint           `json:"origin_airport_id"`
	OriginAirport        Airport        `json:"origin_airport"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
