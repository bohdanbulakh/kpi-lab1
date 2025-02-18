package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	time2 "time"
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

	time := fmt.Sprintf("%v", time2.Now())
	response := map[string]string{"time": time}
	jsonResponse, exception := json.Marshal(response)

	if exception != nil {
		log.Fatal(exception)
	}

	_, exception = writer.Write(jsonResponse)

}
