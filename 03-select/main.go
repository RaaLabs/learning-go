package main

import "fmt"

func main() {

	dataCh := make(chan int)

	// Put things on channel.
	go func() {
		for i := 1; i <= 10; i++ {
			dataCh <- i
		}

		close(dataCh)
	}()

	// Read from channel
	for {
		select {
		case myString, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Printf("contains: %v\n", myString)
		}
	}

}
