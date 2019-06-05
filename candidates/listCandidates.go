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

var database *dynamodb.DynamoDB

type CandidateInfo struct {
	Id         string `json:"id"`
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	Experience int64  `json:"experience"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ListCandidatesResponse struct {
	Candidates []CandidateInfo `json:"candidates"`
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
		database = dynamodb.New(session) // Create DynamoDB client
	}
}

func listCandidates(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	tableName := aws.String(os.Getenv("CANDIDATE_TABLE"))

	// Read all candidates from DynamoDB
	input := &dynamodb.ScanInput{
		TableName: tableName,
	}
	result, _ := database.Scan(input)

	// Construct list of candidates from response
	var candidates []CandidateInfo
	for _, i := range result.Items {
		candidate := CandidateInfo{}
		if err := dynamodbattribute.UnmarshalMap(i, &candidate); err != nil {
			fmt.Println("Failed to unmarshal")
			fmt.Println(err)
		}
		candidates = append(candidates, candidate)
	}

	// Success HTTP response
	body, _ := json.Marshal(&ListCandidatesResponse{
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
