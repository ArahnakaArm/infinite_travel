package routes

import (
	"intravel/controllers"
	"intravel/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoute(app *fiber.App, db *gorm.DB) {
	v1 := app.Group("/v1")
	user := v1.Group("/user")

	user.Post("/register", controllers.NewUserController(db).Register)

	user.Post("/login", controllers.NewUserController(db).Login)

	user.Get("/me", controllers.NewUserController(db).GetUserByMe)

	user.Get("/", controllers.NewUserController(db).GetAllUsers)

	user.Put("/change-password", controllers.NewUserController(db).ChangePassword)

	user.Put("/", controllers.NewUserController(db).UpdateUser)

	user.Patch("/", controllers.NewUserController(db).UpdateSomeFieldUser)

	user.Use("/:id", middleware.AuthConfig, middleware.NewAuthMiddleware(db).CheckAuthFromIdAdmin)
	user.Delete("/:id", controllers.NewUserController(db).DeleteUser)

}
