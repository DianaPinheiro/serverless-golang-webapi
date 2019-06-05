package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Candidate struct {
	Id         string `json:"id"`
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	Experience int64  `json:"experience"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

var ddb *dynamodb.DynamoDB

// In this function we perform some initialization logic like making a database connection to DynamoDB.
// init function is automatically called before main()
func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
	}
}

// addCandidate handler function parses the request body for a candidate information
func addCandidate(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		id        = uuid.Must(uuid.NewV4(), nil).String()
		tableName = aws.String(os.Getenv("CANDIDATE_TABLE"))
	)

	// Initialize Candidate Information
	candidate := &Candidate{
		Id:        id,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	// Parse request body
	_ = json.Unmarshal([]byte(request.Body), candidate)

	// Write to DynamoDB
	item, _ := dynamodbattribute.MarshalMap(candidate)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}

	// Insert a new row to our DynamoDB table
	if _, err := ddb.PutItem(input); err != nil {
		return events.APIGatewayProxyResponse{ // Error HTTP response
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	} else {
		body, _ := json.Marshal(candidate)
		return events.APIGatewayProxyResponse{ // Success HTTP response
			Body:       string(body),
			StatusCode: 200,
		}, nil
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func main() {
	lambda.Start(addCandidate)
}
