package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	TicketId     uint           `gorm:"PRIMARY_KEY;autoIncrement:false" json:"ticket_id"`
	CustomerId   uint           `json:"customer_id"`
	TicketNumber string         `gorm:"not null" json:"ticket_number"`
	FlightId     uint           `json:"flight_id"`
	Flight       Flight         `json:"flight"`
	Seat         string         `gorm:"not null;unique" json:"seat"`
	Status       string         `gorm:"not null" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

/* type UserTest struct {
	MemberNumber string       `gorm:"PRIMARY_KEY" json:"member_number"`
	CreditCards  []CreditCard `gorm:"foreignKey:user_number;references:member_number"`
}

type CreditCard struct {
	gorm.Model
	Number     string
	UserNumber string `json:"user_number"`
}
*/

/* type UserTest struct {
	UserTestId  uint         `gorm:"PRIMARY_KEY"`
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	CreditCardId string
	Number       string
	UserRefer    string
}

*/
/* type UserTest struct {
	MemberNumber uint         `gorm:"PRIMARY_KEY"`
	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number     string
	UserNumber uint
}
*/
/* type UserTest struct {
	gorm.Model
	Name       string     `gorm:"index"`
	CreditCard CreditCard `gorm:"foreignkey:UserName;references:name"`
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}
*/
