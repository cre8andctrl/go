package main

import (
"fmt"
)

func main() {
	/*for i :=0; i<5; i++ {
		fmt.Println(i)
	}*/
	

	/*
	Range
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}*/

	
	//Map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("index", key, "value", value)
	}



	


}