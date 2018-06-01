package main

import (
	"fmt"
 	"strconv"
 	)

func main() {
	x := 5
	fmt.Printf(strconv.Itoa(x))
	fmt.Printf("Hello, world.\n")

	//var a [5]int
	//a[2] = 7

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
}