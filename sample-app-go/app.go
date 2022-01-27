package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Request received")
	msg, exist := os.LookupEnv("HELLO_MSG")
	if exist {
		fmt.Fprintf(w, "<h1>Hello %s!</h1>\n", msg)
	} else {
		fmt.Fprintf(w, "<h1>What are you waiting for?</h1>\n")
	}
}

func main() {
	log.Print("Server started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

