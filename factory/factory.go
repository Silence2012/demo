package main

import "fmt"

type (
	Caculator interface {
		caculate(...int) int
	}
	caculatorImpl struct {
		options CaculatorOptions
	}
	CaculatorOptions struct {
		Basenumber int
	}
)

func NewCaculator(options CaculatorOptions) Caculator {
	return &caculatorImpl{
		options: options,
	}
}
func (c caculatorImpl) caculate(number ...int) (output int) {
	output = c.options.Basenumber
	fmt.Println("output is ", output)
	for i, v := range number {
		fmt.Println("i, v is ", i, v)
		output += i * v
	}
	return
}

func main() {
	var o CaculatorOptions
	o.Basenumber = 3
	l := NewCaculator(o)
	m := 2
	n := 3
	ret := l.caculate(m, n)
	fmt.Println(ret)

	fmt.Println("vim-go")
}
