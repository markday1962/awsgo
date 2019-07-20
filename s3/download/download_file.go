package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	DownloadFromS3Bucket()
}

func DownloadFromS3Bucket() {
	bucket := "s3_bucket"
	item := "a_file"

	file, err := os.Create(item)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Initialize a session in eu-west-1 that the SDK will use to load
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
