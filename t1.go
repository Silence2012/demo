package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "erlove"
	if !strings.Contains(a, "eri3") {
		fmt.Println("true")
	} else {
		fmt.Println("not contain")
	}
}
