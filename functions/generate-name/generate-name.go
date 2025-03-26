package main

import (
	"context"
	"encoding/json"
	"os"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	ArabicName string `json:"arabicName"`
	Type       string `json:"type"`
}

type NameConfig struct {
	Types map[string][]struct {
		Name    string `json:"name"`
		Meaning string `json:"meaning"`
	} `json:"types"`
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var request RequestBody
	if err := json.Unmarshal([]byte(req.Body), &request); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	file, err := os.Open("name-types.json")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	defer file.Close()

	var config NameConfig
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	names, ok := config.Types[request.Type]
	if !ok || len(names) == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	rand.Seed(time.Now().UnixNano())
	selected := names[rand.Intn(len(names))]

	response := map[string]string{
		"chineseName": selected.Name,
		"meaning":     selected.Meaning,
	}

	jsonResponse, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResponse),
	}, nil
}

func main() {
	lambda.Start(handler)
}