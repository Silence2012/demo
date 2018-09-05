package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := `processSpecification=Sp\u00e9cification du processus`
	fmt.Println(input)
	fmt.Println(strconv.Unquote("\"" + input + "\""))
}
