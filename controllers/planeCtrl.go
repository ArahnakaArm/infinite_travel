package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PlaneController interface {
	CreatePlane(c *fiber.Ctx) error
}

type planeController struct {
	db *gorm.DB
}

func NewPlaneController(db *gorm.DB) PlaneController {
	db.AutoMigrate(models.PlaneM{})

	return planeController{db}
}

func (s planeController) CreatePlane(c *fiber.Ctx) error {

	planeReq := models.PlaneM{}

	if err := c.BodyParser(&planeReq); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	var count int64

	s.db.Model(&models.PlaneM{}).Where("plane_name = ?", planeReq.PlaneName).Or("plane_code = ?", planeReq.PlaneCode).Count(&count)

	if count > 0 {
		return services.ConflictResponse(c)
	}

	u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	plane := models.PlaneM{
		PlaneId:   uint(u64),
		AirlineId: planeReq.AirlineId,
		PlaneName: planeReq.PlaneName,
		PlaneCode: planeReq.PlaneCode,
		Status:    planeReq.Status,
		ImgUrl:    planeReq.ImgUrl,
		Model:     planeReq.Model,
	}

	if tx := s.db.Create(&plane); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}
