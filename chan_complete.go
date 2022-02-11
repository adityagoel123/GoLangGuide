// channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("-----")

}
