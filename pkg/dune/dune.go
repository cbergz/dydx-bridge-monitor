package dune

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadToDune(filepath string) {
	apiKey := "azqtH0LzxiCssnnL86uSkjdhDixni70r"
	csvFilePath := filepath
	url := "https://api.dune.com/api/v1/table/upload/csv"

	file, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	headers := map[string]string{
		"X-Dune-Api-Key": apiKey,
	}

	payload := map[string]interface{}{
		"table_name":  "dydx_bridge_events",
		"description": "dYdX bridge events mapped to bech32 addresses",
		"is_private":  false,
		"data":        string(data),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON payload:", err)
		return
	}

	response, err := sendPostRequest(url, headers, jsonPayload)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}

	fmt.Println("Response status code:", response.Status)
	responseData, _ := io.ReadAll(response.Body)
	fmt.Println("Response content:", string(responseData))
}

func sendPostRequest(url string, headers map[string]string, data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}
