package controllers

import (
	"intravel/models"
	"intravel/services"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

type TicketController interface {
	CreateTicket(c *fiber.Ctx) error
}

type ticketController struct {
	db *gorm.DB
}

func NewTicketController(db *gorm.DB) TicketController {
	db.AutoMigrate(models.Ticket{})

	return ticketController{db}
}

func (s ticketController) CreateTicket(c *fiber.Ctx) error {

	uId, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	ticket := models.Ticket{
		TicketId:   uId.String(),
		CustomerId: "162213fb-1a19-4acf-7500-b0e3922b37a4",
	}

	if tx := s.db.Create(&ticket); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return nil
}
