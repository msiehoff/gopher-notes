package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Useful to have nil channels
// will not block when you try to receive from it
// useful with select statement
// can turn off a part of a select statement

func main() {
	timeoutWithNilChannel()
}

func timeoutWithNilChannel() {
	ch := make(chan int)
	go reader(ch)
	go writer(ch)

	time.Sleep(12 * time.Second)
}

func reader(ch chan int) {
	readerT := time.NewTimer(10 * time.Second)
	for {
		select {
			case i := <- ch:
				fmt.Printf("\n message received: %d", i)
			case <- readerT.C:
				// Acts as a timeout
				// once the timer finishes no more messages
				// will be received on ch
				// in the context of a select statement
				// the case above will be ignored
				fmt.Printf("\n readerT stopped")
				return
		}
	}
}

func writer(ch chan int) {
	stopper := time.NewTimer(3 * time.Second)
	restarter := time.NewTimer(6 * time.Second)
	finalStopper := time.NewTimer(10 * time.Second)

	savedCh := ch

	for {
		select {
			case ch <- rand.Intn(100):
			case <- stopper.C:
				fmt.Printf("\npause transmitting")
				ch = nil
			case <- restarter.C:
				fmt.Printf("\nrestart transmitting")
				ch = savedCh
			case <- finalStopper.C:
				fmt.Printf("\nstop transmitting")
				return
		}
	}

}