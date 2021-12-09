package controllers

import (
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

type FlightController interface {
	CreateFlight(c *fiber.Ctx) error
	GetAllFlight(c *fiber.Ctx) error
}

type flightController struct {
	db *gorm.DB
}

func NewFlightController(db *gorm.DB) FlightController {
	db.AutoMigrate(models.Flight{})

	return flightController{db}
}

func (s flightController) CreateFlight(c *fiber.Ctx) error {

	flightReqBody := models.Flight{}

	if err := c.BodyParser(&flightReqBody); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	uId, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	flight := models.Flight{
		FlightId:   uId.String(),
		FlightName: flightReqBody.FlightName,
	}

	if tx := s.db.Create(&flight); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}

func (s flightController) GetAllFlight(c *fiber.Ctx) error {
	offset := -1
	limit := -1

	if c.Query("limit") != "" {
		limitInt, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			return services.MissingAndInvalidResponse(c)
		}

		limit = limitInt
	}

	if c.Query("offset") != "" {
		offsetInt, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			return services.MissingAndInvalidResponse(c)
		}

		offset = offsetInt
	}

	tickets := []models.Ticket{}
	ticketsTotal := []models.Ticket{}

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Find(&tickets); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Find(&ticketsTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, tickets, len(tickets), len(ticketsTotal))
}
