package main

import (
	"fmt"
	// "sort"
	s "github.com/silabig1294/DsAndAlgo/bubble_sort"
	b "github.com/silabig1294/DsAndAlgo/binary_search"
	l "github.com/silabig1294/DsAndAlgo/linear_search"
)

func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	s.Bubblesort(items)
	// sort.Ints(items)
	val := l.Linearsearch(items,45)
	val2 := b.BinarySearch(items,320)
	fmt.Println(items)
	fmt.Println(val)
	fmt.Println(val2)
}
