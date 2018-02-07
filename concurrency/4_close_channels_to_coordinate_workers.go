package main

import (
	"fmt"
	"time"
)

// use closing of a channel as communication

func main() {
	//startWorkAllAtOnce()

	stopSimultaneously()
}

// ------------------ Coordinate work to be started simultaneously ------------------
func startWorkAllAtOnce() {
	// use closing of a channel to start group of worker
	// at the same time (all at once)
	goCh := make(chan bool)
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), goCh)
	}

	// demonstrate workers are waiting on the goCh to close
	time.Sleep(3 * time.Second)

	// Close goCh to kick off all workers simultaneously
	close(goCh)

	// wait so that go routines have a chance to print out
	// each msg before the program terminates
	time.Sleep(3 * time.Second)
}

func printer(msg string, goCh chan bool) {
	// when goCh is closed, this will receive a nil value
	// and the thread will continue
	<- goCh

	fmt.Printf("%s\n", msg)
}

// ------------------ Coordinate work to stop simultaneously ------------------

func stopSimultaneously() {
	stopCh := make(chan bool)

	for i := 0; i < 10; i++ {
		go stopper(fmt.Sprintf("printer: %d", i), stopCh)
	}

	time.Sleep(10 * time.Second)
	close(stopCh)
	time.Sleep(3 * time.Second)
}
func stopper(msg string, stopCh chan bool) {
	for {
		select {
			default:
				fmt.Println(msg)
			case <- stopCh:
				return
		}
	}
}