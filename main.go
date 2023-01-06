package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo"
)

func configS3() *s3.Client {

	creds := credentials.NewStaticCredentialsProvider(os.Getenv("S3_ACCESS_KEY_ID"), os.Getenv("S3_SECRET_ACCESS_KEY"), "")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(os.Getenv("S3_REGION")))

	if err != nil {
		log.Fatal(err)
	}

	return s3.NewFromConfig(cfg)

}

func DeleteImageFromS3(c echo.Context) error {
	e := echo.New()

	e.Start(":")

	awsClient := configS3()
	log.Println("s3 Connected<<<<<<<<<<>>>>>>>>>>>>>", awsClient)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String("mpt-images"),
		Key:    aws.String("pic.png"),
	}

	_, err := awsClient.DeleteObject(context.TODO(), input)

	if err != nil {

		fmt.Println("Got an error deleting item:")
		fmt.Println(err)
		return err

	}

	return e.AcquireContext().JSON(http.StatusOK, "Object Deleted Successfully")

}

func main() {
	Route()

	configS3()

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// post methid using form-data
func Data(c echo.Context) error {
	name := c.FormValue("name")
	return c.String(http.StatusOK, name)
}

//Query parameters

func Params(c echo.Context) error {
	name := c.QueryParam("name")
	log.Println(name)
	return c.String(http.StatusOK, name)

}

func Route() {
	e := echo.New()

	e.GET("/hello", hello)
	e.DELETE("/delete", DeleteImageFromS3)

	e.POST("/name", Data)

	e.GET("/name", Params)
	// add middleware and routes
	// ...
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
