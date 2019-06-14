package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Instance struct {
	Code int32
	Name string
}

func main() {

	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	//in := "i-032b17747be967271"
	var il = []*string{aws.String("i-xxxxxx"), aws.String("i-xxxxx")}
	ec2Svc := ec2.New(sess)

	// Call to get detailed information on each instance
	// result is *ec2.DescribeInstancesOutput
	//input := &ec2.DescribeInstancesInput{
	//	InstanceIds: []*string{
	//		aws.String(in),
	//	},
	//}
	input := &ec2.DescribeInstancesInput{
		InstanceIds: il,
	}
	result, err := ec2Svc.DescribeInstances(input)

	if err != nil {
		fmt.Println("Error", err)
	}

	for idx, _ := range result.Reservations {
		for _, inst := range result.Reservations[idx].Instances {
			if *inst.State.Code == 16 {
				fmt.Printf("The instance %v is running %v\n", *inst.InstanceId, *inst.State.Code)
			} else {
				fmt.Printf("The instance %v is not running, current state is %v\n", *inst.InstanceId, *inst.State.Name)
			}
		}
	}
}
