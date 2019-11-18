package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every Two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)

		select { // select statements allows us to receive from w/e channel is ready
		case msg1 := <-c1:
			fmt.Fprintln(msg1)
		case msg2 := <-c2:
			fmt.Fprintln(msg2)
		}
	}
}
