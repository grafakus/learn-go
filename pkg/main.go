package main

import (
	"fmt"
	"math/cmplx"
	"time"
)

func sum(a, b int) int {
	return a + b
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func split(number int) (p1, p2 int) {
	p1 = number / 10
	p2 = number % 10
	return
}

var bool1 bool
var bool2 bool = true
var bool3 = true

func main() {
	// variables
	fmt.Println("\nVariables:")

	fmt.Println("Hello World!")
	fmt.Println("The time is", time.Now())

	fmt.Println("sum =", sum(7, 35))
	fmt.Println(swap("World", "Hello"))
	fmt.Println(split(128))

	fmt.Println("bool1", bool1)
	bool1 = true
	fmt.Println("bool1", bool1)

	fmt.Println("bool2", bool2)
	fmt.Println("bool3", bool3)

	val := 7
	fmt.Println("val =", val)

	// types
	fmt.Println("\nTypes:")

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("type=%T, value=%v\n", ToBe, ToBe)
	fmt.Printf("type=%T value=%v\n", MaxInt, MaxInt)
	fmt.Printf("type=%T, value=%v\n", z, z)

	// zero values
	fmt.Println("\nZero values:")

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
