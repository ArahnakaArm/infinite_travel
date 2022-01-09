package controllers

import (
	"encoding/json"
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type FlightController interface {
	CreateFlight(c *fiber.Ctx) error
	GetAllFlight(c *fiber.Ctx) error
	GetFlightById(c *fiber.Ctx) error
	UpdateSomeFieldFlight(c *fiber.Ctx) error
	DeleteFlight(c *fiber.Ctx) error
}

type flightController struct {
	db *gorm.DB
}

func NewFlightController(db *gorm.DB) FlightController {
	db.AutoMigrate(models.Flight{})
	db.AutoMigrate(models.Airport{})

	return flightController{db}
}

func (s flightController) CreateFlight(c *fiber.Ctx) error {

	flightReqBody := models.Flight{}

	if err := c.BodyParser(&flightReqBody); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(flightReqBody); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if !services.DateTimeValidate(flightReqBody.ArriveTime) || !services.DateTimeValidate(flightReqBody.DepartTime) {
		return services.MissingAndInvalidResponse(c)
	}

	var conflictNameCount int64

	s.db.Model(&models.Flight{}).Where("flight_name = ?", flightReqBody.FlightName).Count(&conflictNameCount)

	if conflictNameCount > 0 {
		return services.ConflictResponse(c)
	}

	u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	flight := models.Flight{
		FlightId:             uint(u64),
		FlightName:           flightReqBody.FlightName,
		AirlineId:            flightReqBody.AirlineId,
		PlaneMId:             flightReqBody.PlaneMId,
		DestinationAirportId: flightReqBody.DestinationAirportId,
		OriginAirportId:      flightReqBody.OriginAirportId,
		DepartTime:           flightReqBody.DepartTime,
		ArriveTime:           flightReqBody.ArriveTime,
	}

	if tx := s.db.Create(&flight); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}

func (s flightController) GetAllFlight(c *fiber.Ctx) error {
	offset := -1
	limit := -1
	searchQuery := "%%"

	if c.Query("search") != "" {
		searchQuery = "%" + c.Query("search") + "%"
	}

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

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Preload("PlaneM").Preload("Airline").Preload("DestinationAirport").Preload("OriginAirport").Where("flight_name LIKE ? ", searchQuery).Find(&flights); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Where("flight_name LIKE ?", searchQuery).Find(&flightsTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, flights, len(flights), len(flightsTotal))
}

func (s flightController) GetFlightById(c *fiber.Ctx) error {

	flightId := c.Params("id")

	flightRes := models.Flight{}

	if tx := s.db.Preload("PlaneM").Preload("Airline").Preload("DestinationAirport").Preload("OriginAirport").First(&flightRes, "flight_id = ?", flightId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResData(c, flightRes)
}

func (s flightController) UpdateSomeFieldFlight(c *fiber.Ctx) error {

	flightId := c.Params("id")

	var result map[string]interface{}
	json.Unmarshal([]byte(c.Body()), &result)

	if elm, ok := result["flight_name"]; ok {
		var count int64

		s.db.Model(&models.Flight{}).Where("flight_name = ?", elm).Not("flight_id = ?", flightId).Count(&count)

		if count > 0 {
			return services.ConflictResponse(c)
		}
	}

	flight := models.Flight{}

	if tx := s.db.Model(&flight).Where("flight_id = ?", flightId).Updates(result); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	flightRes := models.Flight{}

	if tx := s.db.Preload("PlaneM").Preload("Airline").Preload("DestinationAirport").Preload("OriginAirport").First(&flightRes, "flight_id = ?", flightId); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponseResData(c, flightRes)
}

func (s flightController) DeleteFlight(c *fiber.Ctx) error {

	flightId := c.Params("id")

	flightModel := models.Flight{}

	if tx := s.db.Where("flight_id = ?", flightId).Delete(&flightModel); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponse(c)
}
