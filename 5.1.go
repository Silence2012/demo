package main

import "fmt"

func main() {
	var ch1,ch = make(chan int,1), make(chan int)
	ch2 := make(chan int,1)
	str1 := []string{"1","2","3"}
	str2 := []string{"a","b","c"}
	go func() {
		for k,v := range str1{
			<-ch1
			fmt.Println(v)
			ch2 <- k


		}

	}()
	go func() {

		for k,v := range str2{
			<-ch2
			fmt.Println(v)
			ch1<- k
		}
        ch <- 1
	}()
	ch1 <- 1
	<-ch
}
