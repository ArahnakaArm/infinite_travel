package routes

import (
	"intravel/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CustomerRoute(app *fiber.App, db *gorm.DB) {

	v1 := app.Group("/v1")

	customer := v1.Group("customer")

	customer.Post("/register", controllers.NewCustomerController(db).CreateCustomer)

	customer.Get("/", controllers.NewCustomerController(db).GetAllCustomers)
}
