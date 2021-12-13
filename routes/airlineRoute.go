package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AirlineRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	airline := v1.Group("airline")

	airline.Post("/", controllers.NewAirlineController(db).CreateAirline)

	airline.Get("/", controllers.NewAirlineController(db).GetAllAirline)

}
