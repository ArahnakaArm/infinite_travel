package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TicketRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	ticket := v1.Group("ticket")

	ticket.Post("/", controllers.NewTicketController(db).CreateTicket)

}
