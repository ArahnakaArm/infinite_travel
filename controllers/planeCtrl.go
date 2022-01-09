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

type PlaneController interface {
	CreatePlane(c *fiber.Ctx) error
	GetAllPlane(c *fiber.Ctx) error
	GetPlaneById(c *fiber.Ctx) error
	DeletePlane(c *fiber.Ctx) error
	UpdateSomeFieldPlane(c *fiber.Ctx) error
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

	if errs := validator.Validate(planeReq); errs != nil {
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

func (s planeController) GetAllPlane(c *fiber.Ctx) error {
	offset := -1
	limit := -1
	searchQuery := "%%"
	/* 	planeCodeQuery := "%%" */
	/* 	fmt.Println(string(c.Request().URI().QueryString())) */
	/* planeQuery := map[string]interface{}{"plane_name": "Plane"} */

	/* 	result["name"] = "noob" */
	/* 	values := url.Values{}
	   	values.Add("api_key", "key_from_environment_or_flag")
	   	values.Add("another_thing", "foobar")
	   	query := values.Encode()
	*/
	/* 	fmt.Println(query) */
	/* fmt.Println(planeQuery) */
	/* return nil */

	if c.Query("search") != "" {
		searchQuery = "%" + c.Query("search") + "%"
	}

	planeQuery := map[string]interface{}{}

	if c.Query("airline_id") != "" {
		planeQuery["airline_id"] = c.Query("airline_id")
	}

	/* 	if c.Query("plane_code") != "" {
		planeCodeQuery = "%" + c.Query("plane_code") + "%"
	} */
	/* 	fmt.Println(planeNameQuery) */

	/* 	return nil */

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

	planes := []models.PlaneM{}
	planesTotal := []models.PlaneM{}

	/* 	fmt.Println(planeNameQuery) */

	if tx := s.db.Order("created_at desc").Preload("Airline").Limit(limit).Offset(offset).Where(planeQuery).Where("plane_name LIKE ? OR plane_code LIKE ?", searchQuery, searchQuery).Find(&planes); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Order("created_at desc").Preload("Airline").Where(planeQuery).Where("plane_name LIKE ? OR plane_code LIKE ?", searchQuery, searchQuery).Find(&planesTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, planes, len(planes), len(planesTotal))
}

func (s planeController) GetPlaneById(c *fiber.Ctx) error {

	planeId := c.Params("id")

	plane := models.PlaneM{}

	if tx := s.db.First(&plane, "plane_id = ?", planeId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResData(c, plane)
}

func (s planeController) DeletePlane(c *fiber.Ctx) error {

	planeId := c.Params("id")

	plane := models.PlaneM{}

	if tx := s.db.Where("plane_id = ?", planeId).Delete(&plane); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponse(c)
}

func (s planeController) UpdateSomeFieldPlane(c *fiber.Ctx) error {

	planeId := c.Params("id")

	var result map[string]interface{}
	json.Unmarshal([]byte(c.Body()), &result)

	if elm, ok := result["plane_name"]; ok {
		var count int64

		s.db.Model(&models.PlaneM{}).Where("plane_name = ?", elm).Not("plane_id = ?", planeId).Count(&count)

		if count > 0 {
			return services.ConflictResponse(c)
		}
	}

	if elm, ok := result["plane_code"]; ok {
		var count int64

		s.db.Model(&models.PlaneM{}).Where("plane_code = ?", elm).Not("plane_id = ?", planeId).Count(&count)

		if count > 0 {
			return services.ConflictResponse(c)
		}
	}

	plane := models.PlaneM{}

	if tx := s.db.Model(&plane).Where("plane_id = ?", planeId).Updates(result); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	planeRes := models.PlaneM{}

	if tx := s.db.First(&planeRes, "plane_id = ?", planeId); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponseResData(c, planeRes)

}
