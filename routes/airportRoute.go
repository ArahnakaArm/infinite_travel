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

	airport.Get("/", controllers.NewAirportController(db).GetAllAirport)

	airport.Get("/:id", controllers.NewAirportController(db).GetAirportById)

	airport.Delete("/:id", controllers.NewAirportController(db).DeleteAirport)

	airport.Patch("/:id", controllers.NewAirportController(db).UpdateSomeFieldAirPort)

}
