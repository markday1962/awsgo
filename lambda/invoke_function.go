//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/lambda-go-example-run-function.html
package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"os"
	"strconv"
)

type getItemsResponseBody struct {
	Result string `json:"result"`
	Host   string `json:"host"`
}

type getItemsResponse struct {
	StatusCode int                  `json:"statusCode"`
	Body       getItemsResponseBody `json:"body"`
}

func InvokeFunction() string {
	var region = "eu-west-1"
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region)},
	})
	if err != nil {
		fmt.Println(err)
	}

	input := &lambda.InvokeInput{
		FunctionName: aws.String("aist-get-live-cipher"),
	}

	svc := lambda.New(sess)
	result, err := svc.Invoke(input)
	if err != nil {
		fmt.Println(err)
	}

	var resp getItemsResponse

	err = json.Unmarshal(result.Payload, &resp)
	if err != nil {
		fmt.Println("Error unmarshalling response")
		os.Exit(0)
	}

	// If the status code is NOT 200, the call failed
	if resp.StatusCode != 200 {
		fmt.Println("Error getting livecipher hostname, StatusCode: " + strconv.Itoa(resp.StatusCode))
		os.Exit(0)
	}

	//If the result is failure, we got an error
	if resp.Body.Result == "failure" {
		fmt.Println("Failed to get livecipher hostname")
		os.Exit(0)
	}

	// Print out host name
	if len(resp.Body.Host) == 0 {
		fmt.Println("livecipher hostname not returned")
		os.Exit(0)
	}

	return resp.Body.Host
}

func main() {
	lc := InvokeFunction()
	fmt.Println(lc)
}
