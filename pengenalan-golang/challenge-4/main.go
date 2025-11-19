package main

import "fmt"

func main() {
	data := []int{5, 6, 7, 9, 15, 10, 3, 5, 6, 7, 8, 4, 8, 2, 1, 4, 9, 15, 27, 10, 1}

	start := 6
	end := 18

	hasil := data[start : end+1]

	fmt.Println("Tampilan nilai dari index 6 sampai 18", hasil)
}
