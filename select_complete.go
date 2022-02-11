// select example
package main

import (
	"fmt"
	"time"
)

func main() {

	out := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f\n", val)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}
}
