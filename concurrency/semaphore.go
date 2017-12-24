package main

import(
	"math/rand"
	"fmt"
	"time"
	"strconv"
)

func main() {
	var inputs []string
	numInputs := 100

	for i := 0; i < numInputs; i++ {
		input := rand.Int()

		inputs = append(inputs, strconv.Itoa(input))
	}

	results := doWork(inputs, 10)

	fmt.Printf("\n=== Results ===\n%+v", results)
}

// Semaphore
func doWork(inputs []string, maxWorkers int) []string {
	inputCh := make(chan string)
	outputCh := make(chan string)
	quitCh := make(chan bool)

	defer close(inputCh)
	defer close(outputCh)

	// create limited number of workers to do work
	for i := 0; i < maxWorkers; i++ {
		go worker(i, inputCh, outputCh, quitCh)

	}

	var results []string

	for _, input := range inputs {

		select {
			case inputCh <- input:
				continue
			case workerResult := <- outputCh:
				results = append(results, workerResult)
			default:
				// if no workers are ready to accept the next input
				// & outputCh is not sending anything, sleep for 10ms
				// as to not block
				time.Sleep(10 * time.Millisecond)
		}

		inputCh <- input
		result := <- outputCh
		results = append(results, result)
	}

	close(quitCh)

	return results
}

func worker(workerID int, inputCh, outputCh chan string, quitCh chan bool) {
	for {
		select {
			case input := <- inputCh:
				fmt.Printf("\nWorker %d Received: %s", workerID, input)
				// do work here
				outputCh <- input
			case <- quitCh:
				// all the work has been processed.
				// end the go routine
				fmt.Printf("\nWorker %d received finishedCh", workerID)
				return
			default:
				// if nothing is ready sleep briefly as to not block
				time.Sleep(10 * time.Millisecond)
		}
	}
}

