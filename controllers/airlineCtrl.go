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

type AirlineController interface {
	CreateAirline(c *fiber.Ctx) error
	GetAllAirline(c *fiber.Ctx) error
	GetAirlineById(c *fiber.Ctx) error
	UpdateSomeField(c *fiber.Ctx) error
	DeleteAirline(c *fiber.Ctx) error
}

type airlineController struct {
	db *gorm.DB
}

func NewAirlineController(db *gorm.DB) AirlineController {
	if viper.GetString("ctrl.autoMigrate") == "true" {
		db.AutoMigrate(models.Airline{})
	}

	return airlineController{db}
}

func (s airlineController) CreateAirline(c *fiber.Ctx) error {

	airlineReq := models.Airline{}

	if err := c.BodyParser(&airlineReq); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(airlineReq); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	var count int64

	s.db.Model(&models.Airline{}).Where("airline_name = ?", airlineReq.AirlineName).Or("airline_code = ?", airlineReq.AirlineCode).Count(&count)

	if count > 0 {
		return services.ConflictResponse(c)
	}

	u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	airlineStore := models.Airline{
		AirlineId:   uint(u64),
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

func (s airlineController) GetAllAirline(c *fiber.Ctx) error {
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

	airlines := []models.Airline{}
	airlineTotal := []models.Airline{}

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Preload("PlaneServices").Find(&airlines); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Find(&airlineTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, airlines, len(airlines), len(airlineTotal))

}

func (s airlineController) GetAirlineById(c *fiber.Ctx) error {

	airlineId := c.Params("id")

	airlineRes := models.Airline{}

	if tx := s.db.Preload("PlaneServices").First(&airlineRes, "airline_id = ?", airlineId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResData(c, airlineRes)
}

func (s airlineController) UpdateSomeField(c *fiber.Ctx) error {

	airlineId := c.Params("id")

	var result map[string]interface{}
	json.Unmarshal([]byte(c.Body()), &result)

	if elm, ok := result["airline_name"]; ok {
		var count int64

		s.db.Model(&models.Airline{}).Where("airline_name = ?", elm).Not("airline_id = ?", airlineId).Count(&count)

		if count > 0 {
			return services.ConflictResponse(c)
		}
	}

	if elm, ok := result["airline_code"]; ok {
		var count int64

		s.db.Model(&models.Airline{}).Where("airline_code = ?", elm).Not("airline_id = ?", airlineId).Count(&count)

		if count > 0 {
			return services.ConflictResponse(c)
		}
	}

	airline := models.Airline{}

	if tx := s.db.Model(&airline).Where("airline_id = ?", airlineId).Updates(result); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	airlineRes := models.Airline{}

	if tx := s.db.First(&airlineRes, "airline_id = ?", airlineId); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponseResData(c, airlineRes)
}

func (s airlineController) DeleteAirline(c *fiber.Ctx) error {

	airlineId := c.Params("id")

	airline := models.Airline{}

	if tx := s.db.Where("airline_id = ?", airlineId).Delete(&airline); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponse(c)
}
