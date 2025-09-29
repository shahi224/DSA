// package main

// import "fmt"

// func main() {
// 	var x int = 10
// 	var p *int = &x

// 	fmt.Println("value of x:", x)
// 	fmt.Println("Address of x:", &x)
// 	fmt.Println("value of p:", *p)
// }

// pointer to  a struct
// package main

// import "fmt"

// type person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	p := &person{"shahid", 25}
// 	fmt.Println("Before:", p.Name, p.Age)

// 	p.Age = 26
// 	fmt.Println("After:", p.Name, p.Age)
// }

// pointer in functions

// package main

// import "fmt"

// func UpdateValue(num *int) {
// 	*num = 100
// }

// func main() {
// 	x := 50

// 	UpdateValue(&x)
// 	fmt.Println(x)
// }

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

	slice := make([]int, 3)
	slice[0], slice[1], slice[2] = 1, 2, 3
	fmt.Println("slice using make:", slice)
}
