package main

import "fmt"

func main() {
	var age int = 8 // type can be ommitted because Go can infer variable type from the value, if value type is primitive
	fmt.Println("Age is: ", age)

	var name = "Alphador" // Go inferred because the value type was obvious
	//fmt.Println("Your name is: ", name)
	_ = name

	// Short declaration operator
	s := "Learning golang!"
	fmt.Println(s)
}
