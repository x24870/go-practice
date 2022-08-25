package main

import (
	"fmt"
	"sync"
)

func main() {
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	wg := sync.WaitGroup{}

	pub := func(id int, in chan<- int, done <-chan struct{}) {
		// defer wg.Done()
		for i := 0; i < 10; i++ {
			// select {
			// case <-done:
			// 	return
			// default:
			// }

			select {
			case <-done:
				return
			case in <- i:
				fmt.Printf("pub[%v]: %v\n", id, i)
			}
		}
	}

	sub := func(out <-chan int, done chan<- struct{}) {
		defer wg.Done()
		for i := range out {
			fmt.Println("recv: ", i)
			if i == 5 {
				fmt.Println("recver send DONE")
				// for stopCh sub is the publisher, so it close the stopCh is resonable
				close(done)
				return
			}
		}
	}

	for i := 0; i < 5; i++ {
		// wg.Add(1)
		go pub(i, dataCh, stopCh)
	}

	wg.Add(1)
	go sub(dataCh, stopCh)

	wg.Wait()
}
