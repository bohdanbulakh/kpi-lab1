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

	if err != nil {
		log.Fatal(err)
	}
}

func timeHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	response := map[string]string{"time": time.Now().Format(time.RFC3339)}
	jsonResponse, exception := json.Marshal(response)

	if exception != nil {
		log.Fatal(exception)
	}

	_, exception = writer.Write(jsonResponse)

}
