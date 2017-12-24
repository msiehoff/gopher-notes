# Concurrency

## Basics

**Go Routine:** concurrently running processes/threads. Very lightweight, fast to setup & tear down.

**Channels:** pass messages between go routines. 
- Channels always have a type associated with them, which is the type that is sent/received on the channel.
- Channels are 2 way communication constructs. Anyone with access can both send and receive on them.

```go
// create channel
wordChannel := make(chan string)

// start a go routine to run concurrently
go worker(wordChannel)

// receive single value on channel. 
// Will block until message is received or channel is closed
aWord := <- wordChannel

// (optional) ok boolean return value signifies if the channel is open
anotherWord, ok := <- wordChannel


// receive multiple values on channel (until it's closed)
for word := range wordChannel {
  fmt.Printf("%s", word)
}
```

## Select Keyword

`select`
- enable managing multiple channels simultaneously
- `for` `select` loop is often the core of a go program with many channels (infinite loop waiting for messages from different channels)

```go
// Wait to receive messages on done or word channel
// and process information accordingly
for {
    select {
        case wordChannel <- wordChannel:
            // process message...
        case <- doneChannel:
            // process message...
        }
    }

```

## Channels of type Channel

It is possible to have a channel of type channel. This can be useful if there are many different channels within a program.

```go
// create a channel of string channels
channelCh := make(chan chan string)

// send a message of a new string channel
channelCh <- make(chan string)
```

## Go Routine Semaphore (control how work is processed)

**Purpose:** Control how work is processed by limited number of workers.

**Description:**
Create a number of `worker(inputCh, outputCh chan <type>)` functions.  Each function will receive work to do via `inputCh` & return results via `outputCh`.  Acts as a sort of load balancer. 


## Closing Channels

## Nil Channels

## Buffered Channels

