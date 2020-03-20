package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

//Create a function we use to display errors and exit.
func exitErrorf(msg string, args ...interface{}){
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

//The following example downloads an item from a bucket with the names specified 
//as command line arguments.
func main() {
	if len(os.Args) != 3 {
		exitErrorf("Bucket name and file required\nUsage: %s bucket_name filename", os.Args[0])
	}
	bucket := os.Args[1]
	item := os.Args[2]

	file, err := os.Create(item)
	if err != nil {
		exitErrorf("Unable to create file %q, %v", item, err)
	}

	defer file.Close()

	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config {Region: aws.String("eu-west-1")},)
	
	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key: aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}
	fmt.Println("Downloaded", file.Name, numBytes, "bytes")
}
