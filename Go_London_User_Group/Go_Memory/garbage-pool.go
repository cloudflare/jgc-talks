package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	buffer := make(chan []byte, 50)

	pool := make([][]byte, 200)
	for i := 0; i < 10; i++ {
		go func(offset int) {
			for {
				i := offset+rand.Intn(20)
				
				if pool[i] != nil {
					select {
					case buffer <- pool[i]:
						pool[i] = nil
					default:
					}
				}

				if rand.Intn(10) > 5 {
					var b []byte
					
					select {
					case b = <-buffer:
					default:
						b = makeBuffer()
					}
									
					pool[i] = b
				}
				
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			}
		}(i*20)
	}

	for {
		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc, m.HeapIdle, m.HeapReleased)
		time.Sleep(time.Second)
	}
}
