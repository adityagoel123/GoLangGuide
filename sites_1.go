// Get content type of sites (using channels)
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	// "out" is an channel of type string.

	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")

	// Producing the results onto the Channel.
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Creatinng the Channel of string type.
	ch := make(chan string)

	// Produce Data onto the Channel Basically.
	for _, url := range urls {
		go returnType(url, ch)
	}

	// Consume from Channel Basically.
	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}
