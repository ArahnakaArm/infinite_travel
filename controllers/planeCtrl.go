package controllers

import (
	"intravel/models"
	"intravel/services"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
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

	uId, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	plane := models.PlaneM{
		PlaneId:   uId.String(),
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
