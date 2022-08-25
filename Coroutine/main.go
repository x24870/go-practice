package main

import (
	"fmt"
	"sync"
)

func addBySharedMem(n int) []int {
	arr := []int{}
	var wg sync.WaitGroup

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			arr = append(arr, i)
		}(i)
	}

	wg.Wait()
	return arr
}

func addByLockMem(n int) []int {
	arr := []int{}
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			defer wg.Done()
			arr = append(arr, i)
		}(i)
	}

	wg.Wait()
	return arr
}

func addByChan(n int) []int {
	arr := []int{}
	channel := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(channel chan<- int, i int) {
			channel <- i
		}(channel, i)
	}

	for i := range channel {
		arr = append(arr, i)
		if len(arr) == n {
			break
		}
	}

	close(channel)
	return arr
}

func main() {
	arr := addBySharedMem(10)
	fmt.Println(arr)

	arr = addByLockMem(10)
	fmt.Println(arr)

	arr = addByChan(10)
	fmt.Println(arr)
}
