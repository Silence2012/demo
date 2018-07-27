package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counter int
	wg sync.WaitGroup
	mutex sync.Mutex
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)
	go add(1)
	go add(2)
	fmt.Println(counter)
	wg.Wait()
	fmt.Println(counter)
}

func add (data int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}

        mutex.Unlock()

	}
	fmt.Println(data,counter)
}