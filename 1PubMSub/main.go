package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	pub := func(in chan<- int) {
		for i := 0; i < 20; i++ {
			in <- i
			fmt.Println("send: ", i)
		}
		// close should done by the only publisher
		close(in)

		// wg.Done()
	}

	sub := func(id int, out <-chan int) {
		for i := range out {
			fmt.Printf("recv[%v]: %v\n", id, i)
			time.Sleep(time.Millisecond * time.Duration(500))
		}
		wg.Done()
	}

	// wg.Add(1) // in this case publisher doesn't need to be add to the wait group
	go pub(ch)

	maxRecver := 2
	for i := 0; i < maxRecver; i++ {
		wg.Add(1)
		go sub(i, ch)
	}

	wg.Wait()
}
