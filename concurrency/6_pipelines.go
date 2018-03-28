package main

import "fmt"

// Pipeline: a series of stages of work connected by channels
// Each stage is a group of goroutines running the same function.
//
// In each stage goroutines:
// - receive values from upstream via inbound channels
// - perform some function on that data, usually producing new values
// - send values downstream via outbound channels

// Each stage has any number of inbound and outbound channels, except the first and last stages,
// First stage is sometimes called the source or producer;
// Last stage: called the sink or consumer.


// There is a pattern to our pipeline functions:

// - stages close their outbound channels when all the send operations are done.
// - stages keep receiving values from inbound channels until those channels are closed.

// This pattern allows each receiving stage to be written as a range loop and ensures that
// all goroutines exit once all values have been successfully sent downstream.

// ------------------ Simple Pipeline Example ------------------
func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

// ------------------ First Stage ------------------
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// ------------------ Second Stage ------------------
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Fan Out: Multiple functions reading from the same channel until that channel is closed (distribute work amongst many workers)
// Fan In: Single function reading from multiple inputs and proceeding until all are closed by multiplexing the input channels onto a single channel
// 		   that's closed when all the inputs are closed

func mainWithMerge() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}