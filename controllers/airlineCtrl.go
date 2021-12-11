package controllers

import (
	"intravel/models"
	"intravel/services"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

type AirlineController interface {
	CreateAirline(c *fiber.Ctx) error
}

type airlineController struct {
	db *gorm.DB
}

func NewAirlineController(db *gorm.DB) AirlineController {

	db.AutoMigrate(models.Airline{})

	return airlineController{db}
}

func (s airlineController) CreateAirline(c *fiber.Ctx) error {

	airlineReq := models.Airline{}

	if err := c.BodyParser(&airlineReq); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	var count int64

	s.db.Model(&models.Airline{}).Where("airline_name = ?", airlineReq.AirlineName).Or("airline_code = ?", airlineReq.AirlineCode).Count(&count)

	if count > 0 {
		return services.ConflictResponse(c)
	}

	uId, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	airlineStore := models.Airline{
		AirlineId:   uId.String(),
		AirlineName: airlineReq.AirlineName,
		AirlineCode: airlineReq.AirlineCode,
		Status:      airlineReq.Status,
		ImgUrl:      airlineReq.ImgUrl,
	}

	if tx := s.db.Create(&airlineStore); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)

}
