// Copyright (c) 2014 CloudFlare, Inc.

package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	e0 := l.PushBack(42)
	e1 := l.PushFront(13)
	e2 := l.PushBack(7)
	l.InsertBefore(3, e0)
	l.InsertAfter(196, e1)
	l.InsertAfter(1729, e2)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int))
	}
	fmt.Printf("\n")

	l.MoveToFront(e2)
	l.MoveToBack(e1)
	l.Remove(e0)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int))
	}
	fmt.Printf("\n")
}
