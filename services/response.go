package services

import (
	"intravel/responseMessage"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_SUCCESS,
	})
}

func SuccessResponseResData(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_SUCCESS,
		"resultData":    data,
	})
}

func SuccessResponseResDataRowCount(c *fiber.Ctx, data interface{}, rowCount int, totalCount int) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_SUCCESS,
		"resultData":    data,
		"rowCount":      rowCount,
		"totalCount":    totalCount,
	})
}

func CreatedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusCreated * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_CREATED,
	})
}

func MissingAndInvalidResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusForbidden * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_MISSING_PARAMETER,
	})
}

func InternalErrorResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusInternalServerError * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_INTERNAL_ERROR,
	})
}

func WrongPasswordResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusUnauthorized * 100),
		"resultMessage": responseMessage.RESULT_WRONG_PASSWORD,
	})
}

func NotFoundResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusNotFound * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_DATA_NOT_FOUND,
	})
}

func UnAuthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusUnauthorized * 100),
		"resultMessage": responseMessage.RESULT_UNAUTHORIZED,
	})
}

func ConflictResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusConflict * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_CONFLICT,
	})
}
