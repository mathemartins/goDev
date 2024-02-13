package main

import "fmt"

func main() {
	var a = 2 // Go can infer the variable type
	var b = 3

	c := 3
	d := 5

	fmt.Println(a + b)
	_, _ = c, d // mute the compiler error
}
