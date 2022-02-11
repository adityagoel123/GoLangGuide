// JSON example
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var data = `
{
  "user": "Aditya Goel",
  "type": "deposit",
  "amount": 100000.3
}
`

// Request is a bank transactions
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	// Simulate a file/socket
	rdr := bytes.NewBufferString(data)

	// Getting an Object of Decoder.
	dec := json.NewDecoder(rdr)

	// Creating an Empty Request Object.
	req := &Request{}

	// Trying to decode the Request now.
	err := dec.Decode(req)

	if err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", req)
}
