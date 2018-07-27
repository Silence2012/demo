package main

import (
	"sync"
	"time"
	"fmt"
)

var rw sync.RWMutex
func main() {
	rw.Lock()

	a := "start"
	time.Sleep(20 * time.Second)
	rw.Unlock()


	rw.RLock()
	fmt.Println("1",a)
	time.Sleep(3 * time.Second)
	rw.RUnlock()
	rw.RLock()
	fmt.Println("2",a)
	time.Sleep(5 * time.Second)
	rw.RUnlock()

}