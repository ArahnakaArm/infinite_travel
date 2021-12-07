package controllers

import (
	"intravel/models"
	"intravel/services"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

type CustomerController interface {
	CreateCustomer(c *fiber.Ctx) error
}

type customerController struct {
	db *gorm.DB
}

func NewCustomerController(db *gorm.DB) CustomerController {
	db.AutoMigrate(models.Customer{})

	return customerController{db}
}

func (s customerController) CreateCustomer(c *fiber.Ctx) error {

	customerReq := models.Customer{}

	custModel := models.Customer{}

	if err := c.BodyParser(&customerReq); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	uId, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	tx := s.db.Where("id_card = ? OR visa_number = ? OR mobile_number = ?", customerReq.IdCard, customerReq.VisaNumber, customerReq.MobileNumber).Find(&custModel)

	if tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	if tx.RowsAffected > 0 {
		return services.ConflictResponse(c)
	}

	customer := models.Customer{
		CustomerId:   uId.String(),
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

	return nil
}
