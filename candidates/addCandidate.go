package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	uuid "github.com/satori/go.uuid"
	"os"
	data "serverless-golang-webapi/database"
	models "serverless-golang-webapi/database/Models"
	"time"
)

// addCandidate handler function parses the request body for a candidate information
func addCandidate(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		id        = uuid.Must(uuid.NewV4(), nil).String()
		tableName = aws.String(os.Getenv("CANDIDATE_TABLE"))
	)

	candidate := &models.Candidate{
		Id: id,
		Timestamps: models.Timestamps{
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		},
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
	if _, err := data.DB.PutItem(input); err != nil {
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

func main() {
	lambda.Start(addCandidate)
}
