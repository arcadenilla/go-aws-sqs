package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	credentials, err := sess.Config.Credentials.Get()
	fmt.Println(credentials)

	svc := sqs.New(sess)

	result, err := svc.ListQueues(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success")
	// As these are pointers, printing them out directly would not be useful.
	for i, urls := range result.QueueUrls {
		// Avoid dereferencing a nil pointer.
		if urls == nil {
			continue
		}
		fmt.Printf("%d: %s\n", i, *urls)
	}
}
