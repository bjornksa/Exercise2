// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan bool, number chan<- int) {
	var i = 0
  var finishedChannels = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.
      case v:=<-add_number:
        i += v
        break
      case <-control:
        finishedChannels ++
        break
		}
    if (finishedChannels == 2){
      number <- i
    }
	}
}

func incrementing(add_number chan<-int, finished chan<- bool) {
	for j := 0; j<1000010; j++ {
		add_number <- 1
	}
  finished <- true
	//TODO: signal that the goroutine is finished
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- (-1)
	}
	//TODO: signal that the goroutine is finished
  finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())


	// TODO: Construct the required channels
  finished := make(chan bool)
  add_number := make(chan int)
  number := make(chan int)
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.

	// TODO: Spawn the required goroutines
  go incrementing(add_number, finished)
  go decrementing(add_number, finished)
  go number_server(add_number, finished, number)

	// TODO: block on finished from both "worker" goroutines

	//control<-GetNumber
	Println("The magic number is:", <- number)
	//control<-Exit
}
