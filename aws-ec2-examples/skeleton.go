package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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
	if len(os.Args) != 3 {
		exitErrorf("AMI ID and Instance Type are required"+
			"\nUsage: %s image_id instance_type", os.Args[0])
	}


	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")}, )
	svc := ec2.New(sess)
	if err != nil {
		exitErrorf("Error creating session, %v", err)
	}

}
