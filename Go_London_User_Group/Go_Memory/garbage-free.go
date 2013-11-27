package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
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

		debug.FreeOSMemory()
		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc, m.HeapIdle, m.HeapReleased, makes)
	}
}
