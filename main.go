package deps

import "fmt"

type Dep struct{}

func NewDep() *Dep {
	return &Dep{}
}

func (d *Dep) Version() {
	fmt.Println("Hello, I am v1.1.0")
}
