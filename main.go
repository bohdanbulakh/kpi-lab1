package main

import (
	"fmt"
	"net/http"
)

func main() {
	StartServer(8795)
}

func StartServer(port int) {
	address := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(address, nil)
}
