package main

import (
	"fmt"
)

func generator(strings chan string) {
	strings <- "Five hour's New York jet lag"
	strings <- "and Cayce Pollard wakes in Camden Town"
	strings <- "to the dire and ever-decreasing circles"
	strings <- "of disrupted circadian rhythm."
	close(strings)
}

func main() {
	strings := make(chan string)
	go generator(strings)

	for s := range strings {
		fmt.Printf("%s ", s)
	}
	fmt.Printf("\n");
}
