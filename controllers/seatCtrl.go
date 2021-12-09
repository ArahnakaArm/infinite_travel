package controllers

import (
	"intravel/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SeatController interface {
	CreateSeat(c *fiber.Ctx) error
}

type seatController struct {
	db *gorm.DB
}

func NewSeatController(db *gorm.DB) SeatController {
	db.AutoMigrate(models.Seat{})

	return seatController{db}
}

func (s seatController) CreateSeat(c *fiber.Ctx) error {

	return nil
}
