package main

import (
	"fmt"
	"time"
)

func main() {
	foo := func(done <-chan struct{}) chan struct{} {
		hb := make(chan struct{})
		for {
			select {
			case <-done:
				fmt.Println("foo done")
				close(hb)
				// return
			case hb <- struct{}{}:
				fmt.Println("foo heartbeat")
			case <-time.After(2 * time.Second):
				fmt.Println("foo timeout")
				break
			}
		}

		return hb
	}

	done := make(chan struct{})
	time.AfterFunc(10*time.Second, func() { close(done) })
}
