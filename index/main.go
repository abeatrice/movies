package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var movies = []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}{
	{ID: 1, Name: "Avengers"},
	{ID: 2, Name: "Ant-Man"},
	{ID: 3, Name: "Thor"},
	{ID: 4, Name: "Hulk"},
	{ID: 5, Name: "Iron Man"},
}

func index() (events.APIGatewayProxyResponse, error) {
	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "applicaton/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(index)
}
