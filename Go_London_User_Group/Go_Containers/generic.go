// Copyright (c) 2014 CloudFlare, Inc.

package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushFront("Hello, World!")
	v := l.Front()
	i := v.Value.(int)
	fmt.Printf("Hello, World! == %d\n", i)
}
