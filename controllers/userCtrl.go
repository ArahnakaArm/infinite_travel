package controllers

import (
	"encoding/json"
	"fmt"
	"intravel/models"
	"intravel/services"
	"net/mail"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type UserController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetUserByMe(c *fiber.Ctx) error
	GetAllUsers(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UpdateSomeFieldUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type userController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
	db.AutoMigrate(models.User{})
	return userController{db}
}

////////////////////////////////// Register ///////////////////////////////////////

func (s userController) Register(c *fiber.Ctx) error {

	userRequest := models.RegisterRequest{}

	if err := c.BodyParser(&userRequest); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(userRequest); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if !validateEmail(userRequest.UserName) {
		return services.MissingAndInvalidResponse(c)
	}

	hashedPass, err := hashPassword(userRequest.Password)

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	user := models.User{
		UserName:  userRequest.UserName,
		Password:  hashedPass,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Role:      userRequest.Role,
	}

	s.db.Create(&user)

	return services.CreatedResponse(c)
}

////////////////////////////////// End Register ///////////////////////////////////////

////////////////////////////////// Login ///////////////////////////////////////

func (s userController) Login(c *fiber.Ctx) error {

	loginRequest := models.LoginRequest{}

	if err := c.BodyParser(&loginRequest); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(loginRequest); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	user := models.User{}

	tx := s.db.Where(&models.User{UserName: loginRequest.UserName}).First(&user)

	if tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return services.WrongPasswordResponse(c)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(viper.GetString("appAuth.tokenSecret")))

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
		"token":         t,
	})
}

////////////////////////////////// End Login ///////////////////////////////////////

////////////////////////////////// Get by me  ///////////////////////////////////////

func (s userController) GetUserByMe(c *fiber.Ctx) error {

	splitToken := strings.Split(c.Get("authorization"), "Bearer ")

	if len(splitToken) != 2 {
		return services.MissingAndInvalidResponse(c)
	}

	reqToken := splitToken[1]
	token, err := jwt.Parse(reqToken, nil)
	if token == nil {
		fmt.Println("split error")
		return err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	userId := claims["id"]

	user := models.User{}

	if tx := s.db.First(&user, userId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	fmt.Println(userId)

	fmt.Println(token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
		"resultData":    user,
	})
}

////////////////////////////////// End Get by me  ///////////////////////////////////////

////////////////////////////////// Get All User  ///////////////////////////////////////

func (s userController) GetAllUsers(c *fiber.Ctx) error {

	users := []models.User{}

	result := s.db.Find(&users)

	if result.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
		"resultData":    users,
		"rowCount":      len(users),
	})
}

////////////////////////////////// End Get All User  ///////////////////////////////////////

////////////////////////////////// Change Password  ///////////////////////////////////////

func (s userController) ChangePassword(c *fiber.Ctx) error {

	changePassRequest := models.ChangePasswordRequest{}

	if err := c.BodyParser(&changePassRequest); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(changePassRequest); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	splitToken := strings.Split(c.Get("authorization"), "Bearer ")

	if len(splitToken) != 2 {
		return services.MissingAndInvalidResponse(c)
	}

	reqToken := splitToken[1]
	token, err := jwt.Parse(reqToken, nil)
	if token == nil {
		fmt.Println("split error")
		return err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	userId := claims["id"]

	user := models.User{}

	if tx := s.db.First(&user, userId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePassRequest.OldPassword))
	if err != nil {
		return services.WrongPasswordResponse(c)
	}

	hashedPass, err := hashPassword(changePassRequest.NewPassword)

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	s.db.Model(&user).Update("Password", hashedPass)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
	})
}

////////////////////////////////// End Change Password  ///////////////////////////////////////

////////////////////////////////// Update User  ///////////////////////////////////////

func (s userController) UpdateUser(c *fiber.Ctx) error {

	splitToken := strings.Split(c.Get("authorization"), "Bearer ")

	if len(splitToken) != 2 {
		return services.MissingAndInvalidResponse(c)
	}

	reqToken := splitToken[1]
	token, err := jwt.Parse(reqToken, nil)
	if token == nil {
		fmt.Println("split error")
		return err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	userId := claims["id"]

	updateUser := models.UpdateUser{}

	user := models.User{}

	if err := c.BodyParser(&updateUser); err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	if errs := validator.Validate(updateUser); errs != nil {
		return services.MissingAndInvalidResponse(c)
	}

	var resultMapString map[string]interface{}
	json.Unmarshal([]byte(c.Body()), &resultMapString)

	if tx := s.db.Model(&user).Where("id = ?", userId).Updates(&resultMapString); tx.Error != nil {
		fmt.Println(tx.Error)
		return services.InternalErrorResponse(c)
	}

	if tx := s.db.First(&user, userId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
		"resultData":    user,
	})

}

////////////////////////////////// End Update User  ///////////////////////////////////////

////////////////////////////////// Update Some Field User  ///////////////////////////////////////

func (s userController) UpdateSomeFieldUser(c *fiber.Ctx) error {

	splitToken := strings.Split(c.Get("authorization"), "Bearer ")

	if len(splitToken) != 2 {
		return services.MissingAndInvalidResponse(c)
	}

	reqToken := splitToken[1]
	token, err := jwt.Parse(reqToken, nil)
	if token == nil {
		fmt.Println("split error")
		return err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	userId := claims["id"]

	var result map[string]interface{}
	json.Unmarshal([]byte(c.Body()), &result)

	/* 	fmt.Println(result) */

	user := models.User{}

	if tx := s.db.Model(&user).Where("id = ?", userId).Updates(result); tx.Error != nil {
		fmt.Println(tx.RowsAffected)
		if tx.RowsAffected == 0 {
			return services.MissingAndInvalidResponse(c)
		} else {
			return services.InternalErrorResponse(c)
		}

	}

	if tx := s.db.First(&user, userId); tx.Error != nil {
		return services.NotFoundResponse(c)
	}

	/* s.db.Table("users").Where("id = ?", objId).Updates(result) */

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": "Success",
		"resultData":    user,
	})
}

////////////////////////////////// End Update Some Field User  ///////////////////////////////////////

////////////////////////////////// Delete User  ///////////////////////////////////////

func (s userController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if tx := s.db.Where("id = ?", id).Delete(&models.User{}); tx.Error != nil {
		return services.InternalErrorResponse(c)
	}

	return services.SuccessResponse(c)
}

////////////////////////////////// End Delete User  ///////////////////////////////////////

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPassword(pass string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}
