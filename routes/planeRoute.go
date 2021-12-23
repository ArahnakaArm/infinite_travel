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

	plane.Get("/", controllers.NewPlaneController(db).GetAllPlane)

	plane.Get("/:id", controllers.NewPlaneController(db).GetPlaneById)

	plane.Delete("/:id", controllers.NewPlaneController(db).DeletePlane)

	plane.Patch("/:id", controllers.NewPlaneController(db).UpdateSomeFieldPlane)

}
