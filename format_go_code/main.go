package main

import (
	"fmt"
	"log"
)

const secondsInHour = 3600

func main() {
	fmt.Println() // concatenates using commas and well spaced and adds a \n which is a new line at the end
	fmt.Println("Hello World In Go")
	distance := 60.8 // distance in km
	fmt.Printf("The distance in miles is %f \n", distance*0.621)

	// Go main debugging tool is the log package
	log.Printf("Your age is %d", 21) // doesn't add a new line at the end like Println does
	figure := "Circle"
	radius := 5
	pi := 3.142
	closed := true

	fmt.Printf("The diameter of a %s with radius %d is %f\n", figure, radius, float64(radius)*2*pi)
	fmt.Printf("This is a %q\n", figure)
	fmt.Printf("The diameter of a %v with radius %v is %v\n", figure, radius, float64(radius)*2*pi)
	fmt.Printf("The type of %s is %T\n", figure, figure)
	fmt.Printf("The document file is: %t \n", closed)

	// Base 2
	fmt.Printf("%b \n", 55)

	// Base 16
	fmt.Printf("%x \n", 100)

	// display decimal points
	x := 3.8
	y := 6.7

	fmt.Printf("%.1f multiply by %.1f is %.3f", x, y, x*y)
}
