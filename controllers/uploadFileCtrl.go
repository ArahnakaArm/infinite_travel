package controllers

import (
	"bytes"
	"fmt"
	"intravel/responseMessage"
	"intravel/services"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/spf13/viper"
)

var (
	s3session *s3.S3
)

type UploadFileRes struct {
	Path string `json:"path"`
}

func UploadFile(c *fiber.Ctx) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(viper.GetString("s3.region")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("s3.accessKeyId"), viper.GetString("s3.accessSecretKey"), ""),
	},
	)

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	uploader := s3manager.NewUploader(sess)

	data, err := c.FormFile("file")

	if err != nil {
		return services.MissingAndInvalidResponse(c)
	}

	fileType := data.Header.Get("Content-type")

	fileType = strings.Split(fileType, "/")[1]

	file, err := data.Open()

	if err != nil {
		return err
	}

	defer file.Close()

	fileData, err := ioutil.ReadAll(file)

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	uId, err := uuid.NewV4()

	if err != nil {
		return services.InternalErrorResponse(c)
	}

	s3FolderPath := c.Params("path")

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(viper.GetString("s3.bucketName")),
		Key:    aws.String(fmt.Sprintf("%s/%s.%s", s3FolderPath, uId.String(), fileType)),
		Body:   bytes.NewReader(fileData),
	})

	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	path := &UploadFileRes{
		Path: aws.StringValue(&result.Location),
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"resultCode":    strconv.Itoa(fiber.StatusOK * 100),
		"resultMessage": responseMessage.RESULT_MESSAGE_SUCCESS,
		"resultData":    path,
	})
}
