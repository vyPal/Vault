package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func main() {
	r := gin.Default()

	if gin.Mode() == gin.DebugMode {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	var err error

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS")
	secretAccessKey := os.Getenv("MINIO_SECRET")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"

	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	found, err := minioClient.BucketExists(context.Background(), os.Getenv("MINIO_BUCKET"))
	if err != nil {
		log.Fatalln(err)
	}

	if !found {
		err = minioClient.MakeBucket(context.Background(), os.Getenv("MINIO_BUCKET"), minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}
	
	fmt.Println(minioClient)

	SetupAuth()

	SetupRoutes(r)

	r.Run()
}
