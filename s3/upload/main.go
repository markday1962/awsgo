package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TODO fill these in!
const (
	S3_REGION = "eu-west-1"
	S3_BUCKET = "mark-delete-test"
	ROOT_DIR = "/Users/markday/temp/"
)

func main() {

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	// Get files
	files := ReadFiles()
	for _, file := range files {
		fmt.Println(file)
	}

	// Upload
	err = AddFileToS3(s, files)
	if err != nil {
		log.Fatal(err)
	}
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, files []string) error {

	os.Chdir(ROOT_DIR)
	for _, afile := range files {
		// Open the file for use
		file, err := os.Open(afile)
		if err != nil {
			return err
		}
		defer file.Close()

		// Get file size and read the file content into a buffer
		fileInfo, _ := file.Stat()
		var size int64 = fileInfo.Size()
		buffer := make([]byte, size)
		file.Read(buffer)

		// Config settings: this is where you choose the bucket, filename, content-type etc.
		// of the file you're uploading.

		fmt.Printf("Uploading %s\n", afile)
		_, err = s3.New(s).PutObject(&s3.PutObjectInput{
			Bucket:               aws.String(S3_BUCKET),
			Key:                  aws.String(afile),
			ACL:                  aws.String("private"),
			Body:                 bytes.NewReader(buffer),
			ContentLength:        aws.Int64(size),
			ContentType:          aws.String(http.DetectContentType(buffer)),
			ContentDisposition:   aws.String("attachment"),
			ServerSideEncryption: aws.String("AES256"),
		})
		//return err
	}
	return nil
}

func ReadFiles() []string {
	var fl []string

	files, err := ioutil.ReadDir(ROOT_DIR)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fl = append(fl,file.Name() )
	}

	return fl
}

