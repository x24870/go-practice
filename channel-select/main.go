package main

import (
	"fmt"
	"math/rand"
	"time"
)

func logger(logCh <-chan string, doneChn <-chan struct{}) {
	for {
		select {
		case msg := <-logCh:
			fmt.Println(msg)
		case <-doneChn:
			fmt.Print("Done")
			break
		}
	}
}

func main() {
	logCh := make(chan string)
	doneCh := make(chan struct{})

	go func(n int) {
		max := 10
		min := 1
		// send n times of log to the log channel
		for i := 0; i < n; i++ {
			mSecond := (rand.Intn(max-min) + min) * 100
			time.Sleep(time.Microsecond * time.Duration(mSecond))
			logCh <- fmt.Sprintf("work %v duration %v ms", i, mSecond)
		}

		time.Sleep(5000)
		// send done signal
		// doneCh <- struct{}{}

		// close channels
		close(logCh)
		close(doneCh)
	}(10)

	logger(logCh, doneCh)
}
