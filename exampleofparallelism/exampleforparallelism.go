package main

import (
	"log"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var wg sync.WaitGroup

var threadProfile = pprof.Lookup("threadcreate")

func parallelismInAction() {
	runtime.LockOSThread()
	defer wg.Done()

	time.Sleep(time.Second * 1)

	time.Sleep(time.Second * 1)
	runtime.UnlockOSThread()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	count := 10
	wg.Add(count)

	log.Println("Before thread count:", threadProfile.Count())

	for i := 0; i < count; i++ {
		go parallelismInAction()
	}
	wg.Wait()
	log.Println("After thread count : ", threadProfile.Count())
}
