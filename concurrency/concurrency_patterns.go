package main

/*
----------- Article -----------
https://blog.golang.org/pipelines

----------- Fan in/out -----------

sync.WaitGroup


----------- Stopping Short -----------
There is a pattern to our pipeline functions:

- stages close their outbound channels when all the send operations are done.
- stages keep receiving values from inbound channels until those channels are closed.


----------- Canceling all stages of the pipeline -----------

-

- closing a channel results in all listeners receiving the zero value of the channel type immediately.
- all stages of the pipeline can receive a boolean input channel. When the main method closes the done
	channel each pipeline stage will receive the zero value & can return


----------- Bounded Parallelism -----------

limit number of go routines doing work at a time for a certain stage

----------- Advantages of Pipelines -----------
(as opposed to running each complete slice of work in a separate go routine)

- control how many go routines run at a time at different stages of the pipeline

*/