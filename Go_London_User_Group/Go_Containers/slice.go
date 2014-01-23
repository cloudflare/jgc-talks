// Copyright (c) 2014 CloudFlare, Inc.

package main

import (
	"fmt"
)

type intSlice []int

func (s *intSlice) fill() {
	for i, _ := range *s {
		(*s)[i] = i * 2
	}

	*s = (*s)[1:]
}

func main() {
	primes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	odds := primes[1:7]
	fmt.Printf("%#v\n", odds)
	odds = odds[0 : len(odds)+1]
	copy(odds[4:], odds[3:])
	odds[3] = 9
	fmt.Printf("%#v", odds)

	s := []int{1, 3, 6, 10}
	t := []int{36, 45, 55, 66, 78}

	s = append(s, 15)
	s = append(s, 21, 28)

	s = append(s, t...)

	s = append(s, s...)
	fmt.Printf("%#v\n", s)
}
