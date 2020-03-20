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

//The following example creates a bucket with the name specified as a command line argument.
//You must specify a globally unique name for the bucket.
func main() {
	if len(os.Args) != 2 {
		exitErrorf("Bucket name required\nUsage: %s bucket_name", os.Args[0])
	}
	bucket := os.Args[1]

	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config {Region: aws.String("eu-west-1")},)
	svc := s3.New(sess)

	_, err = svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Unable to delete bucket %q, %v", bucket, err)
	}

	// Wait for the bucket to create
	err = svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		exitErrorf("Error occured while waiting for bucket %q to be deleted, %v", bucket, err)
	}
	fmt.Printf("Bucket %q successfully deleted\n", bucket)
}
