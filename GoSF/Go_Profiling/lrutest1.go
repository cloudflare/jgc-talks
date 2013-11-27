package main

import (
	"bufio"
	"log"
	"net/http"
_	"net/http/pprof"
	"fmt"
	"lru1"
	"os"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6161", nil))
	}()

	cache := lru1.NewCache(1000)
	
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
