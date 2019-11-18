package main

import (
	"fmt"
)

// one channel send and receive
func main() {
	c := make(chan string, 2) // channel of string and buffer channel
	c <- "hello"
	c <- "world"
	// c <- "three"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

	// msg = <-c
	// fmt.Println(msg) // Because buffer is defined at 2, well get deadlock in code and the routine will fail since no one is listening for output of channel
}
