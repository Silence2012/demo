package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main()  {
	syncChan1 := make(chan struct{},1)
	syncChan2 := make(chan struct{},2)
	go func() {
		<-syncChan1
		fmt.Println("f1 Received a sync signal and wait a second ... ")
		time.Sleep(time.Second)
		for {
			if elem,ok := <-strChan; ok {
				fmt.Println("f1 elem is ",elem,ok)
			} else {
				fmt.Println("f1 ",ok)
				break
			}
		}
		fmt.Println("f1 Stopped")
		syncChan2<- struct{}{}
    }()

    go func() {
	    for _, elem := range []string{"a","b","c","d"}{
		    strChan<-elem
		    fmt.Println("f2 Send ",elem)
		    if elem == "c" {
			    syncChan1<- struct{}{}
			    fmt.Println("f2 Sent a sync signal.")
		    }
	    }
	    fmt.Println(" f2 Wait 2 second ... ")
	    time.Sleep(time.Second * 2)
	    close(strChan)
	    syncChan2 <- struct{}{}
	}()

	time.Sleep(30*time.Second)
	//r1 := <-syncChan2
	//fmt.Println("r1 is ",r1)
	//r2 := <-syncChan2
	//fmt.Println("r2 is ",r2)
}