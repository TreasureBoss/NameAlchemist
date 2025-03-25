package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	ArabicName string `json:"arabicName"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	var reqBody RequestBody
	if err := json.Unmarshal(body, &reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"chineseName": fmt.Sprintf("%s的华文名", reqBody.ArabicName),
		"meaning": "美好寓意",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	lambda.Start(Handler)
}