package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := new(Person)
	fmt.Println(p)

	p.Name = "shahid"
	p.Age = 24

	fmt.Println(*p)
}
