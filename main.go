package main

import (
	"fmt"
	"intravel/routes"

	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initTimeZone()
	initConfig()
	app := fiber.New()
	app.Use(cors.New())

	db := initDB()

	loadRoutes(app, db)

	app.Listen(fmt.Sprintf("%s:%s", viper.GetString("app.ip"), viper.GetString("app.port")))
	/* userRepo := repositories.NewUserRepository(db)

	userRepo.CreateUser() */

	/* db.AutoMigrate(User{}) */

	/* fmt.Println("Init Project") */
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	time.Local = ict

}

func initDB() *gorm.DB {

	dsn := fmt.Sprintf("%v:%v@/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"))

	dial := mysql.Open(dsn)

	db, err := gorm.Open(dial)

	if err != nil {
		panic(err)
		/* 	fmt.Println(err) */
	}

	return db

}

func loadRoutes(app *fiber.App, db *gorm.DB) {
	routes.UserRoute(app, db)
	routes.UploadFileRoute(app, db)
	routes.ExploreRoute(app, db)
	routes.CustomerRoute(app, db)
	routes.TicketRoute(app, db)
	routes.FlightRoute(app, db)
	routes.AirlineRoute(app, db)
	routes.PlaneRoute(app, db)
}
