package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type NameConfig struct {
	Types map[string][]struct {
		Name    string `json:"name"`
		Meaning string `json:"meaning"`
	} `json:"types"`
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("文件名: get-name-types.go, 用户IP: %s, 请求方法: %s", req.RequestContext.Identity.SourceIP, req.HTTPMethod)

	// 执行ls -l命令并记录目录信息
	cmd := exec.Command("ls", "-l", "../")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("执行ls命令出错: %v", err)
	} else {
		log.Printf("目录信息:\n%s", string(output))
	}

	file, err := os.Open("../config/name-types.json")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "{\"error\":\"Failed to open config file\"}"}, err
	}
	defer file.Close()

	var config NameConfig
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "{\"error\":\"Failed to parse config file\"}"}, err
	}

	types := make([]string, 0, len(config.Types))
	for k := range config.Types {
		types = append(types, k)
	}

	jsonResponse, _ := json.Marshal(types)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonResponse),
	}, nil
}

func main() {
	lambda.Start(handler)
}
