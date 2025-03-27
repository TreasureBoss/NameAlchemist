package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	ArabicName string `json:"arabicName"`
	Type       string `json:"type"`
}

type NameData struct {
	Name    string `json:"name"`
	Meaning string `json:"meaning"`
}

type NameTypes struct {
	Types map[string][]NameData `json:"types"`
}

func loadNameTypes() (*NameTypes, error) {
	nameTypesFile, err := os.ReadFile("../config/name-types.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read name types file: %v", err)
	}

	var nameTypes NameTypes
	if err := json.Unmarshal(nameTypesFile, &nameTypes); err != nil {
		return nil, fmt.Errorf("failed to parse name types file: %v", err)
	}

	return &nameTypes, nil
}

func validateRequest(request events.APIGatewayProxyRequest) (*RequestBody, error) {
	if request.HTTPMethod != "POST" {
		return nil, fmt.Errorf("method not allowed")
	}

	var reqBody RequestBody
	if err := json.Unmarshal([]byte(request.Body), &reqBody); err != nil {
		return nil, fmt.Errorf("invalid request body: %v", err)
	}

	return &reqBody, nil
}

func generateRandomName(nameTypes *NameTypes, reqBody *RequestBody) (string, string, string) {
	var typeKeys []string
	for key := range nameTypes.Types {
		typeKeys = append(typeKeys, key)
	}

	selectedType := reqBody.Type
	if selectedType == "" {
		selectedType = typeKeys[rand.Intn(len(typeKeys))]
	}

	namesInType := nameTypes.Types[selectedType]
	selectedName := namesInType[rand.Intn(len(namesInType))]

	return selectedType, selectedName.Name, selectedName.Meaning
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("文件名: generate-name.go, 用户IP: %s, 请求方法: %s", request.RequestContext.Identity.SourceIP, request.HTTPMethod)
	reqBody, err := validateRequest(request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf(`{"error":"%s"}`, err),
		}, nil
	}

	nameTypes, err := loadNameTypes()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"error":"%s"}`, err),
		}, nil
	}

	selectedType, name, meaning := generateRandomName(nameTypes, reqBody)

	responseBody, _ := json.Marshal(map[string]string{
		"chineseName": fmt.Sprintf("%s的华文名", reqBody.ArabicName),
		"type":        selectedType,
		"name":        name,
		"meaning":     meaning,
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
