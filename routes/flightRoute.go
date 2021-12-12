package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FlightRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	flight := v1.Group("flight")

	flight.Post("/", controllers.NewFlightController(db).CreateFlight)

	flight.Get("/", controllers.NewFlightController(db).GetAllFlight)

}
