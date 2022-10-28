package main

import (
	"fmt"
	"sync"
)

const semaphore = 3

func main() {
	var wg sync.WaitGroup

	dataCh := make(chan int)

	// Put things on channel.

	wg.Add(1)
	go func() {
		for i := 1; i <= 5; i++ {
			dataCh <- i
		}
		wg.Done()

	}()

	// Put things on channel.
	wg.Add(1)
	go func() {
		for i := 6; i <= 10; i++ {
			dataCh <- i
		}
		wg.Done()

	}()

	done := make(chan struct{})
	// Read from channel
	go func() {

		for myString := range dataCh {
			fmt.Printf("contains: %v\n", myString)
		}
		fmt.Println("Done reading!!!")
		done <- struct{}{}
	}()

	wg.Wait()

	// We can not close the channel from the senders. Instead we need to close the
	// channel when all senders are done. We solve that with waiting for all sending
	// go routines to finish by waiting for the wait group and then closing the channel.
	// The range in the reader will then exit when the chanel is close and send a value
	// on the done channel so we can exit main.
	close(dataCh)
	<-done

}
