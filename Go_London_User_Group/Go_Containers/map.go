// Copyright (c) 2014 CloudFlare, Inc.

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"foo": 42,
		"bar": 54,
		"baz": -1,
	}

	for k := range m {
		fmt.Printf("%s\n", k)
	}
	for _, v := range m {
		fmt.Printf("%d\n", v)
	}
	for k, v := range m {
		fmt.Printf("%s %d\n", k, v)
	}
}
