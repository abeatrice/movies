package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Movie struct {
	ID   string
	Name string
}

func main() {
	movies, err := readMovies("movies.json")
	if err != nil {
		log.Fatal(err)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	for _, movie := range movies {
		fmt.Println("inserting:", movie.Name)
		av, err := dynamodbattribute.MarshalMap(movie)
		if err != nil {
			log.Fatal(err)
		}
		_, err = svc.PutItem(&dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String("movies"),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func readMovies(fileName string) ([]Movie, error) {
	movies := make([]Movie, 0)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return movies, err
	}

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return movies, err
	}

	return movies, nil
}
