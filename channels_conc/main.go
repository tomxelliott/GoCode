package main 

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := make(chan int)
	// can make channels with a length:
	// c := make(chan int, 100)
	// Buffered channels don't add more capacity, they merely provide a queue.
	// They are a good way to deal with a sudden spike.
	for i := 0; i <4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}
	// select ensures we are not using up an infinite amount of memory hoping a worker frees up.
	for {
		select {
		case c <-rand.Int():
			//fmt.Println(len(c))
			//time.Sleep(time.Millisecond * 50)
		case t := <-time.After(time.Millisecond * 100):
			fmt.Println("timed out at", t)
		default:
			// this can be left empty to silently drop the data
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}


type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

// Data can be shared between goroutines safely by using Channels
// To make a channel:

// A channel like everything else in Go has a type
// This is the type of data that will be passing through our newly created channel
/*
to pass a channel to a function, signature looks like:
func worker(c chan int) {
}
*/

// Channels support two operations: receiving and sending.
// Send to a channel by:
// CHANNEL <- DATA 

// Received from a channel by:
// VAR := <-CHANNEL
// The arrow points in the direction that data flows.
// When sending, the data flows into the channel.
// When receiving, the data flows out of the channel.
// In addition there is also blocking - execution is halted.


// If more data coming in than we can handle, we can use Buffered Channels.
