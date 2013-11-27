package main

import (
	"bufio"
	"runtime/pprof"
	"fmt"
	"lru4"
	"os"
	"strings"
)

func main() {
	f, _ := os.Create("lrutest4.cpuprofile")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	cache := lru4.NewCache(1000)

	count := 0
	miss := 0

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		f := strings.Split(in.Text(), "|")
		if len(f) > 2 {
			email := strings.Split(f[2], "@")
			if len(email) > 1 {
				domain := email[1]
				
				if i := cache.Get(domain); i == nil {
					cache.Put(domain, f[2])
					miss += 1
				}
				
				count += 1
			}
		}
	}

	fmt.Printf("%d total %d misses\n", count, miss)
}
