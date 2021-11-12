package repositories

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	UserName  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	LastLogin time.Time `gorm:"default:null"`
}

type UserRepository interface {
	CreateUser() (*user, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&user{})
	return userRepository{db}
}

func (r userRepository) CreateUser() (*user, error) {

	user := user{
		UserName:  "teetawat@gmail.com",
		Password:  "12345678",
		FirstName: "Teetawat",
		LastName:  "Riya",
		Role:      "Admin",
	}

	tx := r.db.Create(&user)

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}

	return &user, nil

}
