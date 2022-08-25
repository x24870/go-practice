package main

import (
	"fmt"
	"sync"
	"time"
)

//Ref: https://ithelp.ithome.com.tw/articles/10218923
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	pub := func(in chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("send: ", i)
			time.Sleep(time.Millisecond * time.Duration(500))
		}
		close(in)
		wg.Done()
	}

	sub := func(out <-chan int) {
		for i := range out {
			fmt.Println("recv: ", i)
		}
		wg.Done()
	}

	wg.Add(1)
	go pub(ch)
	wg.Add(1)
	go sub(ch)

	wg.Wait()
}
