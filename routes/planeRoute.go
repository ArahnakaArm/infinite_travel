package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PlaneRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	plane := v1.Group("plane")

	plane.Post("/", controllers.NewPlaneController(db).CreatePlane)

}
