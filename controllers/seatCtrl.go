package controllers

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type SeatController interface {
	CreateSeat(c *fiber.Ctx) error
}

type seatController struct {
	db *gorm.DB
}

func NewSeatController(db *gorm.DB) SeatController {
	db.AutoMigrate(models.Seat{})

	return seatController{db}
}

func (s seatController) CreateSeat(c *fiber.Ctx) error {

	type seatReqModel struct {
		FligthId   uint `json:"flight_id" validate:"nonzero"`
		MaxRowSeat int  `json:"row_seat" validate:"nonzero"`
	}

	seatBody := seatReqModel{}

	if err := c.BodyParser(&seatBody); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(seatBody); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if seatBody.MaxRowSeat <= 0 {
		return services.MissingAndInvalidResponse(c)
	}

	if seatBody.MaxRowSeat > 10 {
		return services.SeatOverLimitResponse(c)
	}

	if tx := s.db.First(&models.Flight{}, "flight_id = ? ", seatBody.FligthId); tx.Error != nil {
		return services.NotFoundFlightResponse(c)
	}

	var countSeat int64

	if tx := s.db.Model(&models.Seat{}).Where("flight_id = ?", seatBody.FligthId).Count(&countSeat); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	if countSeat > 0 {
		return services.ConflictSeatResponse(c)
	}

	firstSeatIndexs := []string{"A", "B", "C", "D", "E", "F"}

	secondMaxRowSeatIndexs := seatBody.MaxRowSeat

	for i := 0; i < len(firstSeatIndexs); i++ {
		for j := 1; j <= secondMaxRowSeatIndexs; j++ {

			u64, err := strconv.ParseUint(getNumber12digit(), 12, 64)
			if err != nil {
				fmt.Println(err)
			}

			seatCreateModel := models.Seat{
				SeatId:     uint(u64),
				FlightId:   seatBody.FligthId,
				SeatNumber: padNumberWithZero(j) + firstSeatIndexs[i],
			}

			if tx := s.db.Create(&seatCreateModel); tx.Error != nil {
				return services.InternalErrorResponse(c)
			}

			/* fmt.Println(padNumberWithZero(j) + firstSeatIndexs[i]) */
		}
	}

	return services.SuccessResponse(c)
}
