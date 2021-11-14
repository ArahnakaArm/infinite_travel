package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ExploreController interface {
	CreateExplore(c *fiber.Ctx) error
	GetExplore(c *fiber.Ctx) error
	CreateExploreContent(c *fiber.Ctx) error
}

type exploreController struct {
	db *gorm.DB
}

func NewExploreController(db *gorm.DB) ExploreController {
	db.AutoMigrate(models.Explore{})
	db.AutoMigrate(models.ExploreContent{})
	return exploreController{db}
}

func (s exploreController) CreateExplore(c *fiber.Ctx) error {

	exploreRequest := models.ExploreRequest{}

	if err := c.BodyParser(&exploreRequest); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	explore := models.Explore{
		Header: exploreRequest.Header,
	}

	if tx := s.db.Create(&explore); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}

func (s exploreController) CreateExploreContent(c *fiber.Ctx) error {

	exploreContentRequest := models.ExploreContentRequest{}

	if err := c.BodyParser(&exploreContentRequest); err != nil {
		return services.MissingAndInvalidResponse(c)
	}
	intVar, err := strconv.Atoi(c.Params("exploreId"))

	if err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	exploreContent := models.ExploreContent{
		ExploreId: uint(intVar),
		Header:    exploreContentRequest.Header,
		Body:      exploreContentRequest.Body,
		ImageUrl:  exploreContentRequest.ImageUrl,
	}

	if tx := s.db.Create(&exploreContent); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}

func (s exploreController) GetExplore(c *fiber.Ctx) error {

	explore := models.Explore{}

	if tx := s.db.Preload("Content").First(&explore, c.Params("id")); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	fmt.Println(explore)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultData": explore,
	})
}
