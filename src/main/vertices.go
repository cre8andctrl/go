package main

import (
"fmt"
)

func main() {
	

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 2
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

}