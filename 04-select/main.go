package main

import "fmt"

func main() {

	dataCh := make(chan int)
	doneCh := make(chan bool)

	// Put things on channel.
	go func() {
		for i := 1; i <= 10; i++ {
			dataCh <- i
		}

		doneCh <- true
	}()

	// Read from channel
	for {
		select {
		case myString := <-dataCh:
			fmt.Printf("contains: %v\n", myString)
		case <-doneCh:
			return
		}
	}

}
