package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"os"
	data "serverless-golang-webapi/database"
	models "serverless-golang-webapi/database/Models"
)

func listCandidates(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	tableName := aws.String(os.Getenv("CANDIDATE_TABLE"))

	// Read all candidates from DynamoDB
	input := &dynamodb.ScanInput{
		TableName: tableName,
	}
	result, _ := data.DB.Scan(input)

	// Construct list of candidates from response
	var candidates []models.Candidate
	for _, i := range result.Items {
		candidate := models.Candidate{}
		if err := dynamodbattribute.UnmarshalMap(i, &candidate); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		candidates = append(candidates, candidate)
	}

	// Success HTTP response
	body, _ := json.Marshal(&models.ListCandidatesResponse{
		Candidates: candidates,
	})

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil

}

func main() {
	lambda.Start(listCandidates)
}
