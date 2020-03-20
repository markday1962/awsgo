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

	imageId := os.Args[1]
	instanceType := os.Args[2]

	//Initialize the session that the SDK uses to load credentials from the shared credentials file ~/.aws/credentials
	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")}, )
	svc := ec2.New(sess)
	if err != nil {
		exitErrorf("Error creating session, %v", err)
	}

	//Specify the details to create a new instance
	response, err := svc.RunInstances(&ec2.RunInstancesInput{
		// Instance details
		ImageId: aws.String(imageId),
		InstanceType: aws.String(instanceType),
		MinCount: aws.Int64(1),
		MaxCount: aws.Int64(1),
	})

	if err != nil {
		exitErrorf("Could not create instance, %v", err)
	}
	fmt.Printf("Created instance %q\n", *response.Instances[0].InstanceId)
	instance := *response.Instances[0].InstanceId

	//Tag instance
	_, err = svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{response.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key: aws.String("Name"),
				Value: aws.String("Mark Test Instance 2"),
			},
		},
	})
	if err != nil {
		exitErrorf("Error tagging instance %q, %v", instance, err)
	}
	fmt.Printf("Successfully tagged instance %q\n", instance)
}
