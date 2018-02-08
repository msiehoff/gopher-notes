package main

import (
	"time"
	"math/rand"
	"fmt"
	"sync/atomic"
)

// ------------ Unbuffered channels ------------
// - Synchronous communication
//
// Both sender & receiver must be available, otherwise will block execution.
// Sending & receiving happen at the same time.
// If you tried to send on the channel there has to be a go routine listening for anyone to hear the message


// ------------ Buffered Channels ------------
// - Asynchronous communication
//
// - If no go routines are listening at the time a message is sent and there is room in the buffer
// the message is stored in the buffer until a go routine is ready to receive it.
// - Once the buffer is full, the sending threads would have to wait and execution will be blocked for the sender
// - Can be used as a kind of semaphore, to control how many go routines are doing work at a time

func main() {
	semaphoreWork()
}

var (
	workDoneCount int64 = 0
)

func semaphoreWork() {
	// Buffered channel with capacity of 10 messages buffered at a time
	sema := make(chan bool, 20)

	for i := 0; i < 1000; i++ {
		go semaWorker(sema)
	}

	bufferedChannelCapacity := cap(sema)
	for i := 0; i < bufferedChannelCapacity; i++ {
		sema <- true
	}

	time.Sleep(30 * time.Second)
}

func work() {
	atomic.AddInt64(&workDoneCount, 1)
	fmt.Printf("[%d ", workDoneCount)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	atomic.AddInt64(&workDoneCount, -1)
	fmt.Printf("%d] ", workDoneCount)
}

func semaWorker(sema chan bool) {
	<- sema
	work()
	sema <- true
}
