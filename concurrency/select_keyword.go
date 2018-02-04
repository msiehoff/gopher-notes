package main

import (
	"fmt"
	"time"
)

// ------------------- Select Keyword with Channels -------------------
// Normally go programs will communicate on many different channels
// Often one part will need to coordinate between different channels
// listening & transmitting on multiple channels & coordinating when
// certain actions are taken
// e.g. sending input to channel B that was received from channel A


func main() {

	// ----------------- Perform for set number of messages -----------------
	fmt.Printf("\n----------- Set number of messages -----------\n")
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emitWords(wordCh, doneCh)

	fmt.Printf("\n Once upon a time...\n")
	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <- wordCh)
	}

	doneCh <- true
	close(wordCh)


	// ----------------- Perform for set amount of time -----------------
	fmt.Printf("\n----------- Set amount of time -----------\n")
	wordCh2 := make(chan string)

	go emitWordsWithTimer(wordCh2)

	// As long as wordCh2 is open & transmitting words we'll print them
	// Once the channel is closed, move on to the next task
	for word := range wordCh2 {
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\nDone!\n")
}

// Close the channel once we're done transmitting words
func emitWordsWithTimer(wordCh chan string) {
	// Without
	defer close(wordCh)


	words := []string{ "bob", "laba", "law", "blog"}
	index := 0
	timer := time.NewTimer(2 * time.Second)

	for {
		select {
			case wordCh <- words[index]:
				index++
				if index == len(words) {
					index = 0
				}
			case <- timer.C:
				fmt.Printf("\n time's up!\n")
				return
		}
	}
}

func emitWords(wordCh chan string, doneCh chan bool) {
	words := []string{ "bob", "laba", "law", "blog"}

	index := 0

	// Continue sending words to wordCh until sent a message to stop on the doneCh
	// Until this method is told to stop, run an infinite loop monitoring 2 cases:
	// 1. We can send a word on the word channel
	// 2. We are told to stop, by a message on the doneCh
	for {
		select {
			case wordCh <- words[index]:
				index++
				if index == len(words) {
					index = 0
				}
			case <- doneCh:
				close(doneCh)
				return
		}
	}
}
