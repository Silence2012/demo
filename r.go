package main

import (
	"fmt"
	"strings"
)

func main() {
	panicAndRecover()
tags := "hostType:sell;HPGEN9_32C_251G"
strings.Trim()
	strings.Replace(tags, "serverModule:", "", -1)


	fmt.Println("I need to run the statement at any cost!")
}

func panicAndRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("Unable to connect database!")
}
