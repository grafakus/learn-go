package main

import (
	"fmt"
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

var test bool

func main() {
	fmt.Println("Hello World!")
	fmt.Println("The time is", time.Now())

	fmt.Println(sum(7, 35))
	fmt.Println(swap("World", "Hello"))
	fmt.Println(split(128))

	fmt.Println(test)
	test = true
	fmt.Println(test)
}
