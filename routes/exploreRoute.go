package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ExploreRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	explore := v1.Group("explore")

	explore.Post("/", controllers.NewExploreController(db).CreateExplore)

	explore.Post("/content/:exploreId", controllers.NewExploreController(db).CreateExploreContent)

	explore.Get("/:id", controllers.NewExploreController(db).GetExplore)
}
