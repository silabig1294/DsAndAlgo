package main

import (
	"fmt"
	"sort"
	s "github.com/silabig1294/DsAndAlgo/bubble_sort"
	b "github.com/silabig1294/DsAndAlgo/binary_search"
	l "github.com/silabig1294/DsAndAlgo/linear_search"
	g "github.com/silabig1294/DsAndAlgo/generateslices"
	ll "github.com/silabig1294/DsAndAlgo/linked_list"
	// r "github.com/silabig1294/DsAndAlgo/reversing"
)

func main() {

	var i,j int

	fmt.Print("Type two numbers: ")
	fmt.Scanln(&i, &j)
	fmt.Println("Your numbers are:", i, "and", j)
	// var slices []int
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	s.Bubblesort(items)
	slices := g.GenerateSlice(10)
	// s.Bubblesort(slices)
	// a  := r.ReverseInt(slices)
	sort.Slice(slices, func(i, j int) bool {
		return slices[j] < slices[i]
	 })

	// sort.Ints(slices)
	val := l.Linearsearch(items,5)
	val2 := b.BinarySearch(items,320)
	fmt.Println(slices)
	// fmt.Println(a)
	fmt.Println(items)
	fmt.Println(val)
	fmt.Println(val2)

	link := ll.List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)
	link.Insert(88)
	
	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.Head.Key)
	fmt.Printf("Tail: %v\n", link.Tail.Key)
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.Head.Key)
	fmt.Printf("tail: %v\n", link.Tail.Key)
	link.Reverse()
	fmt.Println("\n==============================\n")
}
