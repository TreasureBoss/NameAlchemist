package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	ArabicName string `json:"arabicName"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       `{"error":"Method not allowed"}`,
		}, nil
	}

	var reqBody RequestBody
	if err := json.Unmarshal([]byte(request.Body), &reqBody); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error":"Invalid request body"}`,
		}, nil
	}

	nameTypesFile, err := os.ReadFile("../config/name-types.json")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error":"Failed to read name types file"}`,
		}, nil
	}

	var nameTypesData struct {
		Types map[string][]struct {
			Name    string `json:"name"`
			Meaning string `json:"meaning"`
		} `json:"types"`
	}
	if err := json.Unmarshal(nameTypesFile, &nameTypesData); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error":"Failed to parse name types file"}`,
		}, nil
	}

	// 随机选择一个类型
	var typeKeys []string
	for key := range nameTypesData.Types {
		typeKeys = append(typeKeys, key)
	}
	selectedType := typeKeys[rand.Intn(len(typeKeys))]

	// 从选中的类型中随机选择一个名字
	namesInType := nameTypesData.Types[selectedType]
	selectedName := namesInType[rand.Intn(len(namesInType))]

	responseBody, _ := json.Marshal(map[string]string{
		"chineseName": fmt.Sprintf("%s的华文名", reqBody.ArabicName),
		"type":        selectedType,
		"name":        selectedName.Name,
		"meaning":     selectedName.Meaning,
	})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseBody),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
