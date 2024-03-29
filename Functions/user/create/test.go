package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// args := os.Args[1:]

	// fmt.Printf("Hello %s %s #%s", args[0], args[1], args[2])

	svc := dynamodb.New(session.New(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("aws_access_key_id"), os.Getenv("aws_secret_access_key"), ""),
	}))

	params := &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item: map[string]*dynamodb.AttributeValue{
			"userid":      {S: aws.String("nicolasmartin@gmail.com")},
			"firstname":   {S: aws.String("nicolas")},
			"lastname":    {S: aws.String("martin")},
			"phonenumber": {S: aws.String("11111")},
		},
	}

	resp, err := svc.PutItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		// fmt.Println(err.Error())
		log.Panic(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
