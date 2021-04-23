package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	shawn := Person{
		name: "Shawn",
		age:  24,
	}

	shawn.print()

	shawn.updateName("Sean")
	shawn.print()
}

func (p *Person) updateName(newName string) {
	(*p).name = newName
}

func (p Person) print() {
	// %+v shows keys of structs
	fmt.Printf("%+v\n", p)
}
