package main

import "fmt"

// 4 concurrent workers all pulling items off the jobs que and putting the feed in to the results que at their end
func main() {
	// buffer channels given a size of 100
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results) // 4 concurrent and uses 4 different cpus. We get to take advantage of the multicore processor

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs) // close jobs channel because we have filled up the channel with the input

	for j := 0; j < 100; j++ {
		// After jobs
		fmt.Println(<-results)
	}

}

//           jobs to do    | resuts from channels
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n) // feeds number from job in to results channel
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
