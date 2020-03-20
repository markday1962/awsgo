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

//The following example deletes an item from a bucket with the names specified
//as command line arguments.
func main() {
	if len(os.Args) != 3 {
		exitErrorf("Bucket and object name's are required\nUsage: %s bucket_name object_name", os.Args[0])
	}
	bucket := os.Args[1]
	object := os.Args[2]

	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config {Region: aws.String("eu-west-1")},)

	svc := s3.New(sess)
	_, err = svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(object),
	})
	if err != nil {
		exitErrorf("Unable to delete object %q from bucket %q, %v", object, bucket, err)
	}

	// Wait for the bucket to create
	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(object),
	})
	if err != nil {
		exitErrorf("Error occured while deleting object %q from bucket %q, %v", object, bucket, err)
	}
	fmt.Printf("Object %q successfully deleted from %q\n", object, bucket)
}
