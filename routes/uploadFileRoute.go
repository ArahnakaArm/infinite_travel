package routes

import (
	"intravel/controllers"
	"intravel/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UploadFileRoute(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/v1")
	upload := v1.Group("/upload")
	upload.Post("/:path", middleware.NewAuthMiddleware(db).CheckAuthFromId, controllers.UploadFile)
}
