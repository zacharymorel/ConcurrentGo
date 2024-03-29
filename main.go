package main

import (
	"fmt"
	"time"
)

// NOTE on channels, they communticate and Synchronize

func main() {
	// var waitgroup sync.WaitGroup // this is the counter
	// waitgroup.Add(1)             // I have 1 go routine to wait for.

	c := make(chan string) // make func creates a channel
	go count("Sheep", c)

	for {
		msg, open := <-c // outout from channel
		if !open {
			break
		}
		fmt.Println(msg)
	}

	// go func() {
	// count("Sheep")
	// waitgroup.Done() // decrements the waitground defined counter
	// }()

	// waitgroup.Wait() // Wait, blocks until the counter is 0. IF any go routine have not completed like displayed in func count, then wait will continue to wait.

	// go count("Sheep") // go and run this func in the background and move on to the next line | THIS creates what is called a go routine (go routines run concurrently)
	// count("Fish")
	// NOTE: If you have only a 1 go routine defined, it won't get excuited properly because there won't be any other functions/ process within the main routine

	// time.Sleep(time.Second * 2)
	// fmt.Scanln() // waits of manual user input
	fmt.Println("AFTER GO Routine")
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// fmt.Println(i, thing)
		c <- thing // <- arrow sends a message in to the channel
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // closes a channel when you, as the sender, is done and have sent it's response(s)
}
