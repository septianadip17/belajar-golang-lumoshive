package main

import "fmt"

// function variadic
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// function as paramter
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	// call function variadic
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	// Anonymous functions
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println("Sum:", sum(3, 4))

	// function closure
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2

	// call function as parameter
	result := applyOperation(5, 3, add)
	fmt.Println("Result:", result)

}
