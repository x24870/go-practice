package main

import (
	"fmt"
	"sync"
)

func main() {
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	stopBy := make(chan string)
	wgSub := sync.WaitGroup{}

	// there are multiple publisher and subscriber, so add additional moderator to handle closing data channel
	mod := func(stopBy <-chan string, stopCh chan<- struct{}) {
		s := <-stopBy
		fmt.Printf("MOD: %s", s)
		close(stopCh)
		close(dataCh)
	}

	pub := func(id int, dataCh chan<- int, stopCh <-chan struct{}, stopBy chan<- string) {
		for i := 0; i < 20; i++ {
			// if condition is true, send stop request to moderator
			// use select as a non-block channel writing
			if i >= 7 {
				select {
				case stopBy <- fmt.Sprintf("pub[%v] request close\n", id):
					fmt.Printf("pub[%v] request close\n", id)
				default:
				}
				return
			}

			// check stop channel
			// if moderator tell everybody stop, then stop
			// use select as a non-block channel reading
			select {
			case <-stopCh:
				return
			default:
			}

			// send data, but if there is signal from stopCh then leave
			select {
			case <-stopCh:
				return
			case dataCh <- i:
				fmt.Printf("pub[%v] send: %v\n", id, i)
			}
		}
	}

	sub := func(id int, dataCh <-chan int, stopCh <-chan struct{}, stopBy chan<- string) {
		defer wgSub.Done()
		for {
			// check if moderator request stop
			// use select as a non-block channel reading
			select {
			case <-stopCh:
				return
			default:
			}

			// read data
			select {
			case <-stopCh:
				return
			case i := <-dataCh:
				fmt.Printf("sub[%v] recv: %v\n", id, i)
				// if condition is true, request moderator to close
				if i > 5 {
					// use select as a non-block channel writing
					select {
					case stopBy <- fmt.Sprintf("sub[%v] request to close\n", id):
						fmt.Printf("sub[%v] request to close\n", id)
					default:
					}
					return
				}
			}
		}
	}

	go mod(stopBy, stopCh)

	maxPub := 2
	for i := 0; i < maxPub; i++ {
		go pub(i, dataCh, stopCh, stopBy)
	}

	maxSub := 4
	wgSub.Add(maxSub)
	for i := 0; i < maxSub; i++ {
		go sub(i, dataCh, stopCh, stopBy)
	}

	wgSub.Wait()
}
