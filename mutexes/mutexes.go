// https://gobyexample.com/mutexes
//
// Mutexes - just like RTOS / POSIX / C++ etc
//
// NB. Mutexes MUST NOT be copied, so pass by pointer
//
// We can achieve the same thing using channels instead - way better!
// See stateful-goroutines.go

package main

import (
	"fmt"
	"sync"
)

// This structure holds a map of counters.
// We want to update it concurrently from multiple goroutines, so we've added a Mutex to synchronise it.
// This is a bit weird though. It means the data structure knows about how it'll be used. :(
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex at the start, before accessing counters
// The defer statement causes the Unlock to run at the end
func (c *Container) inc(name string) {
	// Try commenting out these two lines, and running with -race
	// Hurray, it all explodes:
	// ==================
	// WARNING: DATA RACE
	// Read at 0x00c0000a2150 by goroutine 7:
	//   runtime.mapaccess1_faststr()
	//       /snap/go/current/src/runtime/map_faststr.go:13 +0x0
	//   main.(*Container).inc()
	//       /home/kneilb/repos/go-by-example/mutexes.go:29 +0x88
	//   main.main.func1()
	//       /home/kneilb/repos/go-by-example/mutexes.go:43 +0x58
	//   main.main.func2()
	//       /home/kneilb/repos/go-by-example/mutexes.go:50 +0x4d

	// Previous write at 0x00c0000a2150 by goroutine 9:
	//   runtime.mapassign_faststr()
	//       /snap/go/current/src/runtime/map_faststr.go:203 +0x0
	//   main.(*Container).inc()
	//       /home/kneilb/repos/go-by-example/mutexes.go:29 +0xba
	//   main.main.func1()
	//       /home/kneilb/repos/go-by-example/mutexes.go:43 +0x58
	//   main.main.func4()
	//       /home/kneilb/repos/go-by-example/mutexes.go:52 +0x4d
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// The zero value of a mutex is usable, so it doesn't need initialising.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Increment a named counter in a loop, then indicate done on the WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently.
	// All 3 access the same Container (therefore map), 2 access the same key!
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Wait for the goroutines to finish, print out the results
	wg.Wait()
	fmt.Println(c.counters)
}
