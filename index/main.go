package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Movie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func index(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	size, err := strconv.Atoi(request.Headers["Count"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       err.Error(),
		}, nil
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	req, resp := svc.ScanRequest(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Limit:     aws.Int64(int64(size)),
	})
	err = req.Send()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while scanning DynamoDB",
		}, nil
	}
	movies := make([]Movie, 0)
	for _, item := range resp.Items {
		movies = append(movies, Movie{
			ID:   *item["ID"].S,
			Name: *item["Name"].S,
		})
	}
	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error while marshal resp items",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(index)
}
