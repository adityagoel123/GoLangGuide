// channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// close to signal we're done
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

	fmt.Println("-----")

}
