package main

import "fmt"

func gen(nums ...int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, i := range nums {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

// output the square of input int
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	for n := range sq(sq(gen(1, 2, 3))) {
		fmt.Println(n)
	}
}
