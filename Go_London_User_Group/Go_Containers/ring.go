// Copyright (c) 2014 CloudFlare, Inc.

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	parus := []string{"major", "holsti", "carpi"}

	r := ring.New(len(parus))
	for i := 0; i < r.Len(); i++ {
		r.Value = parus[i]
		r = r.Next()
	}

	r.Do(func(x interface{}) {
		fmt.Printf("Parus %s\n", x.(string))
	})
}
