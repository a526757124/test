package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 6, 8, 1, 2, 0}
	sort.Ints(a)

	for _, v := range a {
		fmt.Println(v)
	}
}
