package main

import (
	"fmt"
)

// ------------------ Go Routines ------------------
// Purpose: Run functions in a separate, concurrently processed thread
// very light-weight, can start as many as necessary

// ------------------ Channels ------------------
// Purpose: pass messages between go routines
//
// - Each channel has a type (the type of the message sent on the channel)
//
// - Always thread safe. Only 1 receive can happen at a time.  In the example
// of makeID below, each `uniqueID := <- idChan` would be guaranteed to get
// a unique number.
//
// - If no go routines are available to send on a channel (meaning it wasn't closed
// when it should have been) the program will raise a deadlock exception
//
// - When waiting to receive a message from a channel, the process will block, unless
// the receive statement is in a select
//
// - Channels can be nested, creating a channel who's type is chan string
// var channelCh chan chan string // Send & receive string channels via channelCh





func main() {
	// ------------ Create a new channel ------------
	wordChannel := make(chan string)

	// ------------ Start Go Routine ------------
	// run the emit function in a go routine, concurrently from
	// can run as many go routines at once as you want, since they're
	// small & efficient
	go emit(wordChannel)

	// ------------ Receive from channel - Single ------------
	// receive a single message from a channel
	// this blocks until a message is received unless the channel is closed
	singleWord, isChannelOpen := <- wordChannel

	if !isChannelOpen {
		// False means the channel has been closed
		panic("channel was closed prematurely!")
	}

	fmt.Printf("%s ", singleWord)

	// ------------ Receive until channel closes ------------
	// receive each message successively on the channel
	// until the channel is closed
	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")

	idChan := make(chan int)
	// defer close(idChan)

	go makeID(idChan)

	fmt.Printf("\n==== unique ID's ====\n")
	fmt.Printf("\n%v", <- idChan)
	fmt.Printf("\n%v", <- idChan)
	fmt.Printf("\n%v", <- idChan)
	fmt.Printf("\n%v", <- idChan)
	fmt.Printf("\n%v", <- idChan)
}

// Run infinite loop which sends integers down a channel
// when someone requests a new ID.

func makeID(c chan int) {
	id := 0
	for {
		c <- id
		id++
	}
}

func emit(c chan string) {
	wordsList := []string { "the", "quick", "brown", "fox"}

	for _, word := range wordsList {
		c <- word
	}

	// Channels must be closed
	// If they are not, a fatal exception will occur
	close(c)
}