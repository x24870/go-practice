package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"a", "b", "c"}

	fmt.Println(sort.SearchStrings(s, "b"))
}
