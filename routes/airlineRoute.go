package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AirlineRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	customer := v1.Group("airline")

	customer.Post("/", controllers.NewAirlineController(db).CreateAirline)

}
