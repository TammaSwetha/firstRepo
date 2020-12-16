package main

import "fmt"

func ChangeValue(str *string) {
	*str = "changed!"
}
func ChangeValue2(str string) {
	str = "changed!"
}

func main() {
	// i, j := 10, 20

	// p := &i
	// fmt.Println(*p)

	// *p = 30
	// fmt.Println(i) // i will be 30

	// p = &j
	// fmt.Println(&p) // addaress of j

	// *p = *p * 10 // 20 * 10
	// fmt.Println(j)

	x := 10
	y := &x
	fmt.Println(x, y)
	*y = 20
	fmt.Println(x, y)
	*y = 30
	fmt.Println(x, y)

	fmt.Println("...........")

	toChange := "Hello"
	fmt.Println(toChange)

	ChangeValue(&toChange) // call the functions
	fmt.Println(toChange)

	// another example
	Name := "Swetha"
	var pointer *string = &Name
	fmt.Println(pointer, &pointer) // printed pointer to the pointer first pointer points to the name, second pointer to the &Name

	// ChangeValue2(toChange) // call the functions
	// fmt.Println(toChange)
}
