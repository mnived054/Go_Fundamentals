package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j + "job done"
	}
}

func main() {
	numJobs := 5

	jobsList := []string{"job1", "job2", "job3", "job4", "job5"}

	jobs := make(chan string, numJobs)

	results := make(chan string, numJobs)

	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	for _, job := range jobsList {
		jobs <- job
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
