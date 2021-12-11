package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
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
	db.AutoMigrate(models.Ticket{})
	db.AutoMigrate(models.Seat{})

	return ticketController{db}
}

func (s ticketController) CreateTicket(c *fiber.Ctx) error {

	ticketReqBody := models.Ticket{}

	if err := c.BodyParser(&ticketReqBody); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	uId, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
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

	ticket := models.Ticket{
		TicketId:     uId.String(),
		CustomerId:   ticketReqBody.CustomerId,
		FlightId:     ticketReqBody.FlightId,
		Status:       ticketReqBody.Status,
		TicketNumber: flightTicket,
		Seat:         seat,
	}

	if tx := s.db.Create(&ticket); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	uIdSeat, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	seatInsert := models.Seat{
		SeatId:     uIdSeat.String(),
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

func getTicketNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}

	return string(codes[:])
}

func getSeatNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [2]byte
	for i := 0; i < 2; i++ {
		codes[i] = uint8(48 + r.Intn(5))
	}

	return string(codes[:])
}

func generateTicket(flightName string) string {

	return fmt.Sprintf("%s-%s", flightName, getTicketNumber())
}

func generateSeat() string {

	randomSeatChar := 'A' + rune(rand.Intn(6))

	return fmt.Sprintf("%s-%s", string(randomSeatChar), getSeatNumber())
}
