package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

//Create a function we use to display errors and exit.
func exitErrorf(msg string, args ...interface{}){
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

//The following example lists the items in the bucket with the name specified
//as a command line argument.
func main() {

	if len(os.Args) != 2 {
		exitErrorf("Bucket name required\nUsage: %s bucket_name", os.Args[0])
	}
	bucket := os.Args[1]

	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config {Region: aws.String("eu-west-1")},)

	//Create S3 service client
	svc := s3.New(sess)

	response, err := svc.ListObjectsV2(&s3.ListObjectsV2Input {Bucket: aws.String(bucket)})
	if err != nil {
		exitErrorf("Unable to list objects inbucket %q, %v", bucket, err)
	}

	for _, item := range response.Contents {
		fmt.Println("Name				", *item.Key)
		fmt.Println("Last Modified		", *item.LastModified)
		fmt.Println("Size				", *item.Size)
		fmt.Println("Storage Class		", *item.StorageClass)
		fmt.Println("")
	}


}