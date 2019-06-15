package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"log"
)

func ListFunctions() {
	var region = "eu-west-1"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	})
	svc := lambda.New(sess)

	input := &lambda.ListFunctionsInput{
		//Marker:   aws.String(""),
		MaxItems: aws.Int64(123),
	}

	result, err := svc.ListFunctions(input)
	if err != nil {
		log.Println(err)
	}

	for _, v := range result.Functions {
		log.Printf("%v\n", *v.FunctionName)
	}

}

func main() {
	ListFunctions()
}
