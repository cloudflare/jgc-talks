package main

import (
	"container/list"
	"net"
	"time"
)

func recycler(give, get chan []byte) {
	q := new(list.List)

	for {
		if q.Len() == 0 {
			q.PushFront(make([]byte, 100))
		}

		e := q.Front()

		select {
		case s := <-give:
			q.PushFront(s[:0])

		case get <- e.Value.([]byte):
			q.Remove(e)
		}
	}
}

func main() {
	give := make(chan []byte)
	get := make(chan []byte)

	go recycler(give, get)

	b0 := <- get
	b1 := <- get
	// ... use b0, b1
	give <- b0
	give <- b1
	
	time.Sleep(time.Minute)
}
