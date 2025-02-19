package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	StartServer(8795)
}

func StartServer(port int) {
	http.HandleFunc("/time", timeHandler)

	address := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(address, nil)

	ProcessException(err)
}

func timeHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	response := map[string]string{"time": getTime()}
	jsonResponse, exception := json.Marshal(response)

	ProcessException(exception)

	_, exception = writer.Write(jsonResponse)

	ProcessException(exception)
}

func ProcessException(exception error) {
	if exception != nil {
		log.Fatal(exception)
	}
}

func getTime() string {
	return time.Now().Format(time.RFC3339)
}
