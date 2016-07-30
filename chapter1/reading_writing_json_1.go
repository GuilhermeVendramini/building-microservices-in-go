package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloWorldResponse struct {
	Message string
}

func main() {
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Starting server on port %v\n", 8080)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}