package main

import (
	"fmt"
	"net/http"
	"runtime"
)

type job struct {
	url string
	resp chan *http.Response
}

type worker struct {
	jobs chan *job
	count int
}

func (w *worker) getter(done chan *worker) {
	for {
		j := <- w.jobs
		resp, _ := http.Get(j.url)
		j.resp <- resp
		done <- w
	}
}

func balancer(count int, depth int) chan *job {
	jobs := make(chan *job)
	done := make(chan *worker)
	workers := make([]*worker, count)

	for i := 0; i < count; i++ {
		workers[i] = &worker{make(chan *job, depth), 0}
		go workers[i].getter(done)
	}

	go func() {
		for {
			var free *worker
			min := depth
			fmt.Printf("WORKERS: ")
			for _, w := range workers {
				fmt.Printf("%d ", w.count)
				if w.count < min {
					free = w
					min = w.count
				}
			}
			fmt.Printf("\n")

			var jobsource chan *job
			if free != nil {
				jobsource = jobs
			}

			select {
			case j := <- jobsource:
				free.jobs <- j
				free.count++
				
			case w := <- done:
				w.count--
			}
		}
	}()

	return jobs
}

func get(jobs chan *job, url string, answer chan string) {
	resp := make(chan *http.Response)
	jobs <- &job{url, resp}
	r := <- resp
	if r != nil {
		answer <- r.Request.URL.String()
	} else {
		answer <- url
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	jobs := balancer(10, 10)
	answer := make(chan string)

	for {
		var url string
		if _, err := fmt.Scanln(&url); err != nil {
			break
		}

		go get(jobs, url, answer)
	}

	for _ = range answer {
//		fmt.Printf("%s\n", u)
	}
}
