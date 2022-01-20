package controllers

import (
	"encoding/json"
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type AirportController interface {
	CreateAirport(c *fiber.Ctx) error
	GetAllAirport(c *fiber.Ctx) error
	GetAirportById(c *fiber.Ctx) error
	DeleteAirport(c *fiber.Ctx) error
	UpdateSomeFieldAirPort(c *fiber.Ctx) error
}

type airportController struct {
	db *gorm.DB
}

func NewAirportController(db *gorm.DB) AirportController {

	if viper.GetString("ctrl.autoMigrate") == "true" {
		db.AutoMigrate(models.Airline{})
	}

	return airportController{db}
}

func (s airportController) CreateAirport(c *fiber.Ctx) error {

	airportReq := models.Airport{}

	if err := c.BodyParser(&airportReq); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(airportReq); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	var count int64

	s.db.Model(&models.Airline{}).Where("airport_name = ?", airportReq.AirportName).Or("airport_code = ?", airportReq.AirportCode).Count(&count)

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

func (s airportController) GetAllAirport(c *fiber.Ctx) error {

	offset := -1
	limit := -1

	excludeBody := map[string]interface{}{}

	if c.Query("exclude") != "" {
		excludeIdInt, err := strconv.Atoi(c.Query("exclude"))
		if err != nil {
			return services.InternalErrorResponse(c)
		}
		excludeBody["airport_id"] = excludeIdInt
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

	airports := []models.Airport{}
	airportsTotal := []models.Airport{}

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Not(excludeBody).Find(&airports); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Not(excludeBody).Find(&airportsTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, airports, len(airports), len(airportsTotal))

}

func (s airportController) GetAirportById(c *fiber.Ctx) error {

	airportId := c.Params("id")

	airportRes := models.Airport{}

	if tx := s.db.First(&airportRes, "airport_id = ?", airportId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResData(c, airportRes)
}

func (s airportController) DeleteAirport(c *fiber.Ctx) error {

	airportId := c.Params("id")

	airport := models.Airport{}

	if tx := s.db.Where("airport_id = ?", airportId).Delete(&airport); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}
	return services.SuccessResponse(c)
}

func (s airportController) UpdateSomeFieldAirPort(c *fiber.Ctx) error {

	airportId := c.Params("id")

	var resultBody map[string]interface{}

	json.Unmarshal([]byte(c.Body()), &resultBody)

	if elm, ok := resultBody["airport_name"]; ok {
		var count int64

		s.db.Model(&models.Airport{}).Where("airport_name = ?", elm).Not("airport_id = ?", airportId).Count(&count)

		if count > 0 {
			return services.ConflictResponse(c)
		}
	}

	if elm, ok := resultBody["airport_code"]; ok {
		var count int64

		s.db.Model(&models.Airport{}).Where("airport_code = ?", elm).Not("airport_id = ?", airportId).Count(&count)

		if count > 0 {
			return services.ConflictResponse(c)
		}
	}

	airport := models.Airport{}

	if tx := s.db.Model(&airport).Where("airport_id = ?", airportId).Updates(resultBody); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	airportRes := models.Airport{}

	if tx := s.db.First(&airportRes, "airport_id = ?", airportId); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponseResData(c, airportRes)

}
