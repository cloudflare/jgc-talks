package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
_   "net/http/pprof"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6161", nil))
	}()

	pool := make([][]byte, 20)

	var m runtime.MemStats
	makes := 0
	for {
		b := makeBuffer()
		i := rand.Intn(len(pool))
		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc, m.HeapIdle, m.HeapReleased, makes)
	}
}
