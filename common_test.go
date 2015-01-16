// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import "testing"
import "log"
import "time"

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	j := <-jobs
	log.Println("worker", id, "processing job", j)
	time.Sleep(time.Second)
	results <- j * 2
}

func TestPool(t *testing.T) {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int)
	results := make(chan int)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 9 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= 10; j++ {
		//activeJobCount := len(jobs)
		jobs <- j
	}
	close(jobs)

}
