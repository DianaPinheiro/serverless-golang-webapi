package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"os"
	data "serverless-golang-webapi/database"
	models "serverless-golang-webapi/database/Models"
)

// Get interview candidate by ID sent by request
func getCandidateByID(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var (
		tableName = aws.String(os.Getenv("CANDIDATE_TABLE"))
		// Parse ID from request body
		candidateID = request.PathParameters["id"]
	)

	// Read candidate by ID
	result, err := data.DB.GetItem(&dynamodb.GetItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(candidateID),
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	candidate := &models.Candidate{}
	err = dynamodbattribute.UnmarshalMap(result.Item, candidate)

	if candidate.Id == "" {
		body, _ := json.Marshal("Could not find the candidate with ID " + candidateID)
		return events.APIGatewayProxyResponse{
			Body:       string(body),
			StatusCode: 200,
		}, nil
	}

	body, _ := json.Marshal(candidate)

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(getCandidateByID)
}
