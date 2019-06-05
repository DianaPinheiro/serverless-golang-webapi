package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"os"
)

var db *dynamodb.DynamoDB

type CandidateInformation struct {
	Id         string `json:"id"`
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	Experience int64  `json:"experience"`
}

// In this function we perform some initialization logic like making a database connection to DynamoDB.
// init function is automatically called before main()
func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		db = dynamodb.New(session) // Create DynamoDB client
	}
}

func getCandidateByID(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		tableName = aws.String(os.Getenv("CANDIDATE_TABLE"))
		// Parse ID from request body
		candidateID = request.PathParameters["id"]
	)

	// Read candidate by ID
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(candidateID),
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	candidate := CandidateInformation{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &candidate)

	body, _ := json.Marshal(candidate)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(getCandidateByID)
}
