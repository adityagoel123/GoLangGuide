package main

import (
	"fmt"
	//"os"
)

func main() {
	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println("The value @ this non-existent Index in Slice received is : ", v)
	fmt.Println("We have succesfully recovered from the Erroneous Situation.")
}

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}
