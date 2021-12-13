package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AirportRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	airport := v1.Group("airport")

	airport.Post("/", controllers.NewAirportController(db).CreateAirport)

}
