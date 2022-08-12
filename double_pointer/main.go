package main

import "fmt"

func main() {

	var a int = 64

	fmt.Println("variable decalred before the first pointer :", a)

	fmt.Println("The address of a", &a)

	fmt.Println()

	var Pt1 *int = &a

	// it will give the address of the Pt1
	fmt.Println("variable address decalred of the first pointer Pt1:", Pt1)

	// it will give the value of address of pt1 which has been fetch from a
	fmt.Println("value of Pt1:", *Pt1)

	fmt.Println()

	// the value been changed in the pointer
	*Pt1 = 759

	fmt.Println("the address will be same as the previous pointer value", Pt1)
	fmt.Println("the value will be changed from the previous pointer value Pt1")
	fmt.Println("value of Pt1 30:", *Pt1)

	// taking pointer to pointer to Pt1
	// storing to address of Pt1 into Pt2
	var Pt2 **int = &Pt1

	fmt.Println()

	fmt.Println("the address has been changed from the previous pointer value", Pt2)
	fmt.Println("value of Pt2 39:", **Pt2)

	**Pt2 = 83700

	fmt.Println("the value will be changed from the previous pointer value Pt1")
	fmt.Println("value of Pt2:", **Pt2)

}
