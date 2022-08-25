package main

import (
	"fmt"
	"sync"
)

type Set struct {
	m    map[int]struct{}
	len  int
	lock sync.RWMutex
}

func NewSet() *Set {
	// m := make(map[int]struct{})
	m := map[int]struct{}{}
	// var m map[int]struct{}
	return &Set{
		// m: map[int]struct{}{},
		m: m,
	}
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Add(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[item] = struct{}{}
	s.len = s.Len()
}

func (s *Set) Remove(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.m[item]; !ok {
		return
	}
	delete(s.m, item)
	s.len = s.Len()
}

func (s *Set) Has(item int) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

func (s *Set) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

func (s *Set) List() []int {
	s.lock.Lock()
	defer s.lock.Unlock()
	lst := make([]int, 0, s.len)
	for k := range s.m {
		lst = append(lst, k)
	}

	return lst
}

func main() {
	s := NewSet()
	s.Add(5)
	fmt.Println(s.List())
	s.Add(6)
	fmt.Println(s.List())
	s.Add(6)
	fmt.Println(s.List())
	s.Remove(6)
	fmt.Println(s.List())

	fmt.Println(s.Has(5))
	fmt.Println(s.Has(4))
	fmt.Println(s.IsEmpty())
	s.Clear()
	fmt.Println(s.IsEmpty())
}
