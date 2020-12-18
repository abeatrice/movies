package main

import (
	"encoding/json"
	"net/http"
	"os"

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

func show(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	req, resp := svc.GetItemRequest(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(request.PathParameters["id"]),
			},
		},
	})
	err := req.Send()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "error while fetching from dynamoDB",
		}, nil
	}
	movie := &Movie{
		ID:   *resp.Item["ID"].S,
		Name: *resp.Item["Name"].S,
	}
	json, err := json.Marshal(movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "json marshal error",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(json),
	}, nil
}

func main() {
	lambda.Start(show)
}
