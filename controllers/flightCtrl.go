package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
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

	u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	flight := models.Flight{
		FlightId:   uint(u64),
		FlightName: flightReqBody.FlightName,
		AirlineId:  flightReqBody.AirlineId,
		PlaneMId:   flightReqBody.PlaneMId,
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

	flights := []models.Flight{}
	flightsTotal := []models.Flight{}

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Preload("PlaneM").Preload("Airline").Find(&flights); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Find(&flightsTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, flights, len(flights), len(flightsTotal))
}
