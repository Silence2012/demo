package main

import "fmt"

func A(pdl ...map[string]string) {
	fmt.Println(pdl)

}
func main() {
	r := make(map[string]string)
	r["pdl"] = "iaas"
	r["link"] = "web"
	A(r)
	println("After this, panic will start")
	panic("Panic occoured!")
	println("This line will not appear")
}
