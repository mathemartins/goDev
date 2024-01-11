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

	// redeclaration cannot be done using the short declaration operator unless one variable is new
	car, price := "Aston Martin", 5000
	fmt.Println(car, price)

	// now to redeclare
	car, year := "BMW", 2023 // this worked because year is a new variable
	fmt.Println(car, year)

	// another example
	var opened bool = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int // they were initialized to zero, so they are not new variables, therefore cannot use short declaration
	i, j = 5, 8
	i, j = j, i

	fmt.Println(i, j)

	sum := 5 + 3.5
	fmt.Println(sum)
}
