package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	s    []int
	len  int
	lock sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		s:   []int{},
		len: 0,
	}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) IsEmpty() bool {
	if q.Len() == 0 {
		return true
	}
	return false
}

func (q *Queue) Poll() (int, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.IsEmpty() {
		return 0, false
	}
	tmp := q.s[0]
	q.s = q.s[1:]
	q.len = len(q.s)
	return tmp, true
}

func (q *Queue) Offer(item int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.s = append(q.s, item)
	q.len += 1
}

func main() {
	q := NewQueue()
	fmt.Println(q.IsEmpty(), q.Len(), q)

	q.Offer(1)
	q.Offer(3)
	q.Offer(5)
	fmt.Println(q.IsEmpty(), q.Len(), q)

	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	fmt.Println(q.IsEmpty(), q.Len(), q)

	fmt.Println(q.Poll())
}
