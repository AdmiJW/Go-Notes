package go_by_example

import (
	"fmt"
	"time"
)

// Here, we implement a worker pool using goroutines and channels.

// Here's a worker goroutine.
// We'll run several concurrent instances. These workers will receive work on the `jobs`
// channel and send the corresponding results on `results`.
// We'll sleep a second per job to simulate an expensive task.
func worker2(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerPools() {
	// In order to use our pool of workers, we need to send them work and collect their results.
	// We make 2 channels for this.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	// Here, we send 5 jobs and then close that channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally, we collect all the results of the work.
	// This also ensures that the worker goroutines have finished.
	for a := 1; a <= numJobs; a++ {
		fmt.Println("result obtained:", <-results)
	}
}