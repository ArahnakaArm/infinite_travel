package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TicketController interface {
	CreateTicket(c *fiber.Ctx) error
	GetAllTickets(c *fiber.Ctx) error
}

type ticketController struct {
	db *gorm.DB
}

func NewTicketController(db *gorm.DB) TicketController {

	/* db.AutoMigrate(models.UserTest{})
	db.AutoMigrate(models.CreditCard{})
	*/
	db.AutoMigrate(models.Ticket{})
	db.AutoMigrate(models.Seat{})

	return ticketController{db}
}

func (s ticketController) CreateTicket(c *fiber.Ctx) error {

	ticketReqBody := models.Ticket{}

	if err := c.BodyParser(&ticketReqBody); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	flight := models.Flight{}

	if tx := s.db.Find(&flight, "flight_id = ?", ticketReqBody.FlightId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	flightName := flight.FlightName

	flightTicket := generateTicket(flightName)

	seat := generateSeat()

	/* fmt.Println(seat)
	 */
	/* 	return nil */

	/* 	if seatCount > 0  */

	////// Mock

	//////

	u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	ticket := models.Ticket{
		TicketId:     uint(u64),
		CustomerId:   ticketReqBody.CustomerId,
		FlightId:     ticketReqBody.FlightId,
		Status:       ticketReqBody.Status,
		TicketNumber: flightTicket,
		Seat:         seat,
	}

	if tx := s.db.Create(&ticket); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	u64Seat, err := strconv.ParseUint(getNumber12digit(), 12, 64)
	if err != nil {
		fmt.Println(err)
	}

	seatInsert := models.Seat{
		SeatId:     uint(u64Seat),
		SeatNumber: seat,
		FlightId:   ticketReqBody.FlightId,
	}

	if tx := s.db.Create(&seatInsert); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.CreatedResponse(c)
}

func (s ticketController) GetAllTickets(c *fiber.Ctx) error {
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

	tickets := []models.Ticket{}
	ticketsTotal := []models.Ticket{}

	if tx := s.db.Order("created_at desc").Limit(limit).Offset(offset).Preload("Flight").Find(&tickets); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	if tx := s.db.Find(&ticketsTotal); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return services.SuccessResponseResDataRowCount(c, tickets, len(tickets), len(ticketsTotal))
}
