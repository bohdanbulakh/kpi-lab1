package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	StartServer(8795)
}

func StartServer(port int) {
	address := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
