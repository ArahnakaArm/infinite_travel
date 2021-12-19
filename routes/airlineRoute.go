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

	airline.Get("/:id", controllers.NewAirlineController(db).GetAirlineById)

	airline.Patch("/:id", controllers.NewAirlineController(db).UpdateSomeField)

	airline.Delete("/:id", controllers.NewAirlineController(db).DeleteAirline)
}
