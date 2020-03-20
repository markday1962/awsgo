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

//The following example copies an item from one bucket to another with the names specified
//as command line arguments.
func main() {
	if len(os.Args) != 4 {
		exitErrorf("Source bucket, Object, Target bucket are required" +
			"\nUsage: %s source_bucket object_name target_bucket", os.Args[0])
	}
	sourceb := os.Args[1]
	object := os.Args[2]
	targetb := os.Args[3]
	source := sourceb + "/" + object

	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config {Region: aws.String("eu-west-1")},)
	svc := s3.New(sess)

	_, err = svc.CopyObject(&s3.CopyObjectInput{
		Bucket: aws.String(targetb),
		CopySource: aws.String(source),
		Key: aws.String(object),
	})
	if err != nil{
		exitErrorf("Unable to copy object %q from %q to %q, v%\n", object, sourceb, targetb, err)
	}

	// Wait to see if the object has copied successfully
	err = svc.WaitUntilObjectExists(&s3.HeadObjectInput{
		Bucket: aws.String(targetb),
		Key: aws.String(object),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for object %q to copied from %q to %q, %v", object, sourceb, targetb, err)
	}
	fmt.Printf("Object %q successfully copied from bucket %q to bucket %q\n", object, sourceb, targetb)
}
