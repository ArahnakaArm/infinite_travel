package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AirportController interface {
	CreateAirport(c *fiber.Ctx) error
}

type airportController struct {
	db *gorm.DB
}

func NewAirportController(db *gorm.DB) AirportController {

	db.AutoMigrate(models.Airline{})

	return airportController{db}
}

func (s airportController) CreateAirport(c *fiber.Ctx) error {

	airportReq := models.Airport{}

	if err := c.BodyParser(&airportReq); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	var count int64

	s.db.Model(&models.Airline{}).Where("airline_name = ?", airportReq.AirportName).Or("airline_code = ?", airportReq.AirportCode).Count(&count)

	if count > 0 {
		return services.ConflictResponse(c)
	}

	u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	airportStore := models.Airport{
		AirportId:   uint(u64),
		AirportName: airportReq.AirportName,
		AirportCode: airportReq.AirportCode,
	}

	if tx := s.db.Create(&airportStore); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)

}
