package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SeatRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	seat := v1.Group("seat")

	seat.Post("/", controllers.NewSeatController(db).CreateSeat)

}
