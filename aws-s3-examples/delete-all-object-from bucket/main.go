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

//The following example deletes all the items from a bucket with the bucket name
//specified as a command line argument.
func main() {
	if len(os.Args) != 2 {
		exitErrorf("Bucket name is required\nUsage: %s bucket_name", os.Args[0])
	}
	bucket := os.Args[1]

	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, _ := session.NewSession(&aws.Config {Region: aws.String("eu-west-1")},)
	svc := s3.New(sess)

	//Create a list iterator to iterate through the list of bucket objects and deleting each object.
	iter := s3manager.NewDeleteListIterator(svc, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})

	if err := s3manager.NewBatchDeleteWithClient(svc).Delete(aws.BackgroundContext(), iter); err != nil {
		exitErrorf("Unable to delete objects from bucket %q, %v", bucket, err)
	}
	fmt.Printf("Objects successfully deleted from %q\n", bucket)
}
