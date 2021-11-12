package models

import (
	"time"
)

type User struct {
	CommonModelFields
	UserName   string    `gorm:"not null" json:"user_name"`
	Password   string    `gorm:"not null" json:"password"`
	FirstName  string    `gorm:"not null" json:"first_name"`
	LastName   string    `gorm:"not null" json:"last_name"`
	Role       string    `gorm:"not null" json:"role"`
	ProfileUrl string    `gorm:"default:null" json:"profile_url"`
	LastLogin  time.Time `gorm:"default:null" json:"last_login"`
}

type UpdateUser struct {
	FirstName string `gorm:"not null" json:"first_name" validate:"nonnil,nonzero"`
	LastName  string `gorm:"not null" json:"last_name" validate:"nonnil,nonzero"`
	Role      string `gorm:"not null" json:"role" validate:"nonnil,nonzero"`
}

type RegisterRequest struct {
	UserName  string `json:"user_name" validate:"nonnil,nonzero"`
	Password  string `json:"password" validate:"nonnil,nonzero"`
	FirstName string `json:"first_name" validate:"nonnil,nonzero"`
	LastName  string `json:"last_name" validate:"nonnil,nonzero"`
	Role      string `json:"role" validate:"nonnil,nonzero"`
}

type LoginRequest struct {
	UserName string `json:"user_name" validate:"nonnil,nonzero"`
	Password string `json:"password" validate:"nonnil,nonzero"`
}

type ChangePasswordRequest struct {
	OldPassword string `bson:"oldPassword" json:"old_password" validate:"nonnil,nonzero,min=8"`
	NewPassword string `bson:"newPassword" json:"new_password" validate:"nonnil,nonzero,min=8"`
}
