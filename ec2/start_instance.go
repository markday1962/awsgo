package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

var in = "i-0b82a82e9a2142729"
var region = "eu-west-1"

//setting our session to eu-west-1
func GetSession() session.Session {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	})

	if err != nil {
		fmt.Printf("Error caught creating a session: %v", err)
	}
	return *sess
}

func StartInstance(s *session.Session){
	log.Printf("Request to start instance %v has sent!\n", in)
	//input for our start ec2 instance request
	svc := ec2.New(s)
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(in),
		},
	}
	log.Printf("Instance input values: %v", input)

	result, err := svc.StartInstances(input)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Response from starting %v :%v\n",in,result)
	log.Println("StartInstance has been finished!")
}

func main(){
	s := GetSession()
	StartInstance(&s)
}
