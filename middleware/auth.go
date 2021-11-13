package middleware

import (
	"fmt"
	"intravel/models"
	"intravel/services"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
)

type AuthMiddleware interface {
	CheckAuthFromId(c *fiber.Ctx) error
	CheckAuthFromIdAdmin(c *fiber.Ctx) error
}

type authMiddleware struct {
	db *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) AuthMiddleware {
	return authMiddleware{db}
}

var AuthConfig = jwtware.New(jwtware.Config{
	SigningMethod:  "HS256",
	SigningKey:     []byte("secret"),
	SuccessHandler: SuccessValidate,
	ErrorHandler:   FailAuth,
})

func SuccessValidate(c *fiber.Ctx) error {
	return c.Next()
}

func FailAuth(c *fiber.Ctx, e error) error {
	return services.UnAuthorizedResponse(c)
}

func (m authMiddleware) CheckAuthFromId(c *fiber.Ctx) error {

	splitToken := strings.Split(c.Get("authorization"), "Bearer ")
	reqToken := splitToken[1]
	token, _ := jwt.Parse(reqToken, nil)

	if token == nil {
		return services.UnAuthorizedResponse(c)
	}

	/* if err != nil {
		fmt.Println(err)
		return services.UnAuthorizedResponse(c)
	}
	*/
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["id"]

	user := models.User{}
	if tx := m.db.First(&user, id); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return c.Next()

}

func (m authMiddleware) CheckAuthFromIdAdmin(c *fiber.Ctx) error {
	fmt.Println("CheckAuth")
	splitToken := strings.Split(c.Get("authorization"), "Bearer ")
	if len(splitToken) != 2 {
		return services.MissingAndInvalidResponse(c)
	}

	reqToken := splitToken[1]
	token, _ := jwt.Parse(reqToken, nil)
	if token == nil {
		return services.UnAuthorizedResponse(c)
	}

	/* 	if err != nil {
	   		return services.UnAuthorizedResponse(c)
	   	}
	*/
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["id"]

	user := models.User{}

	if tx := m.db.First(&user, id); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	if user.Role != "Admin" {
		return services.UnAuthorizedResponse(c)
	}

	fmt.Println(user.Role)
	return c.Next()

}
