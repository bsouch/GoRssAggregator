package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func jsonResponse(writer http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal Json response: %v", payload)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	writer.Write(data)
}

func errorJsonResponse(writer http.ResponseWriter, statusCode int, errorMessage string) {
	if statusCode >= 500 {
		log.Printf("Responding with status code %v error:", errorMessage)
	}

	responseObj := errorResponse{Error: errorMessage}
	jsonResponse(writer, statusCode, responseObj)
}

func jsonDeserialise(r *http.Request, params interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&params)
}
