package main

import (
	"fmt"
)

func main() {

	x, y, z := 42, "James Bond", true

	fmt.Println(x,y,z)
	fmt.Printf("%v %T\n", x,x)
	fmt.Printf("%v %T\n", y,y)
	fmt.Printf("%v %T\n", z,z)
}