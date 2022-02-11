// Makign HTTP calls
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Job is a job description
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// POST request
	job := &Job{
		User:   "Aditya",
		Action: "Stellar Affection",
		Count:  11,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: can't encode job - %s", err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
