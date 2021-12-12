package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CustomerController interface {
	CreateCustomer(c *fiber.Ctx) error
	GetAllCustomers(c *fiber.Ctx) error
}

type customerController struct {
	db *gorm.DB
}

func NewCustomerController(db *gorm.DB) CustomerController {
	db.AutoMigrate(models.Customer{})

	return customerController{db}
}

/////////////////////////////// Create Customer ///////////////////////////////

func (s customerController) CreateCustomer(c *fiber.Ctx) error {

	customerReq := models.Customer{}

	custModel := models.Customer{}

	if err := c.BodyParser(&customerReq); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if !validateEmail(customerReq.UserName) {
		return services.MissingAndInvalidResponse(c)
	}

	hashedPass, err := hashPassword(customerReq.Password)

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	tx := s.db.Where("id_card = ? OR visa_number = ? OR mobile_number = ? OR user_name = ?", customerReq.IdCard, customerReq.VisaNumber, customerReq.MobileNumber, customerReq.UserName).Find(&custModel)

	if tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	if tx.RowsAffected > 0 {
		return services.ConflictResponse(c)
	}

	u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	customer := models.Customer{
		CustomerId:   uint(u64),
		UserName:     customerReq.UserName,
		Password:     hashedPass,
		FirstName:    customerReq.FirstName,
		LastName:     customerReq.LastName,
		MiddleName:   customerReq.MiddleName,
		IdCard:       customerReq.IdCard,
		VisaNumber:   customerReq.VisaNumber,
		MobileNumber: customerReq.MobileNumber,
		Nation:       customerReq.Nation,
		Gender:       customerReq.Gender,
	}

	if tx := s.db.Create(&customer); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}

/////////////////////////////// End Create Customer ///////////////////////////////

/////////////////////////////// Get Customers //////////////////////////////////////

func (s customerController) GetAllCustomers(c *fiber.Ctx) error {
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

	customers := []models.Customer{}
	customersTotal := []models.Customer{}

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Preload("Tickets").Preload("Tickets.Flight").Find(&customers); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Find(&customersTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, customers, len(customers), len(customersTotal))
}

/////////////////////////////// End Get Customers //////////////////////////////////////
