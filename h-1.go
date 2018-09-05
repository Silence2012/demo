package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "jim", " a name here")
}
func main() {
	flag.Parse()
	fmt.Println("good job ", name)

}
