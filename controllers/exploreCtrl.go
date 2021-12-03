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
	GetExploreById(c *fiber.Ctx) error
	CreateExploreContent(c *fiber.Ctx) error
	GetExplores(c *fiber.Ctx) error
	DeleteExplore(c *fiber.Ctx) error
}

type exploreController struct {
	db *gorm.DB
}

func NewExploreController(db *gorm.DB) ExploreController {
	db.AutoMigrate(models.Explore{})
	db.AutoMigrate(models.ExploreContent{})

	db.Migrator().DropColumn(&models.ExploreContent{}, "body")
	db.Migrator().DropColumn(&models.ExploreContent{}, "header")

	return exploreController{db}
}

//////////////////////// Create Explore ///////////////////////////////////

func (s exploreController) CreateExplore(c *fiber.Ctx) error {

	exploreRequest := models.ExploreRequest{}

	if err := c.BodyParser(&exploreRequest); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	explore := models.Explore{
		Title:     exploreRequest.Title,
		Author:    exploreRequest.Author,
		Paragraph: exploreRequest.Paragraph,
		ImageUrl:  exploreRequest.ImageUrl,
	}

	if tx := s.db.Create(&explore); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}

//////////////////////// End Create Explore ///////////////////////////////////

//////////////////////// Create Explore Content ///////////////////////////////////

func (s exploreController) CreateExploreContent(c *fiber.Ctx) error {

	exploreContentRequest := models.ExploreContentRequest{}

	if err := c.BodyParser(&exploreContentRequest); err != nil {
		return services.MissingAndInvalidResponse(c)
	}
	intVarId, err := strconv.Atoi(c.Params("exploreId"))

	if err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	exploreContent := models.ExploreContent{
		ExploreId: uint(intVarId),
		Title:     exploreContentRequest.Title,
		Paragraph: exploreContentRequest.Paragraph,
		ImageUrl:  exploreContentRequest.ImageUrl,
	}

	if tx := s.db.Create(&exploreContent); tx.Error != nil {
		if tx.RowsAffected == 0 {
			return services.NotFoundResponse(c)
		} else {
			return services.InternalErrorResponse(c)
		}

	}

	return services.CreatedResponse(c)
}

//////////////////////// End Create Explore Content ///////////////////////////////////

//////////////////////// Get Explore By id ///////////////////////////////////

func (s exploreController) GetExploreById(c *fiber.Ctx) error {

	explore := models.Explore{}

	if tx := s.db.Preload("Content").First(&explore, c.Params("id")); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	fmt.Println(explore)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
		"resultData":    explore,
	})
}

//////////////////////// End Get Explore By id ///////////////////////////////////

//////////////////////// Get Explore By id ///////////////////////////////////

func (s exploreController) GetExplores(c *fiber.Ctx) error {
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

	explores := []models.Explore{}

	exploresTotal := []models.Explore{}

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Preload("Content").Find(&explores); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Find(&exploresTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
		"resultData":    explores,
		"rowCount":      len(explores),
		"totalCount":    len(exploresTotal),
	})
}

//////////////////////// End Get Explore By id ///////////////////////////////////

////////////////////////////////// Delete User  ///////////////////////////////////////

func (s exploreController) DeleteExplore(c *fiber.Ctx) error {
	id := c.Params("id")
	if tx := s.db.Where("explore_id = ?", id).Delete(&models.Explore{}); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	if tx := s.db.Where("explore_id = ?", id).Delete(&models.ExploreContent{}); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponse(c)
}

////////////////////////////////// End Delete User  ///////////////////////////////////////
