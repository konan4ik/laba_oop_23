package observer

import "fmt"

type doctor struct {
	name string
}

func (d doctor) Update(action string) {
	fmt.Println(action)
}

func NewDoctor(name string) *doctor {
	return &doctor{name: name}
}