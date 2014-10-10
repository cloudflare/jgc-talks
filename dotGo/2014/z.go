// Code from my dotGo.eu 2014 presentation
//
// Copyright (c) 2014 John Graham-Cumming
//
// Implement a factory and a task. Call run() on your factory.

package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)

type task interface {
	process()
	print()
}

type factory interface {
	make(line string) task
}

func run(f factory) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.make(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.print()
	}
}

func main() {
	// run(&myFactory{})
}
