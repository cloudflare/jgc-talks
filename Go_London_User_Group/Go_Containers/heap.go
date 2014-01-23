// Copyright (c) 2014 CloudFlare, Inc.

package main

import (
	"container/heap"
	"fmt"
)

type OrderedInts []int

func (h OrderedInts) Len() int { return len(h) }
func (h OrderedInts) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h OrderedInts) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *OrderedInts) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *OrderedInts) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

func main() {
	h := &OrderedInts{33, 76, 55, 24, 48, 63, 86, 83, 83, 12}
	heap.Init(h)
	fmt.Printf("min: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
