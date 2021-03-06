package main

import (
	"log"
	"time"
	"github.com/eugenewyj/go-sample/goinaction/ch07/work"
	"sync"
)

// names provides a set of names todisplay
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Task implements the Worker interface.
func (m *namePrinter) Task()  {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// main is the entry point for all Go program.
func main() {
	// Create a work pool with 2 goroutines.
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(2 * len(names))

	for i := 0; i < 2; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specific name.
			np := namePrinter{
				name: name,
			}

			go func() {
				// Submit the task to be worked on. When RunTask
				// returns we konw it is being handled.
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed.
	p.Shutdown()
}
