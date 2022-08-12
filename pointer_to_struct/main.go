package main

import "fmt"

type Employee struct {

	// taking variable
	Name  string
	Empid int
}

func main() {

	// creating an instatnce of the Employee stuct type

	Emp := Employee{"ABC", 1234}

	// Here, it is the pointer to the struct
	Pts := &Emp //PointerAddress

	fmt.Println(Pts)

	// accessing the struct fields(liem employee's name)
	// using a pointer but here we are not using derefrencing explicitly

	fmt.Println(Pts.Name)

	// same as above by explicitely using
	// dereferencing concept
	// means the result will be the same

	fmt.Println((*Pts).Name)

	// we can also modify the value of the structure members or structure
	// literal by using the pointer

	Pts.Name = "shubham"

	fmt.Println(Pts)

}
