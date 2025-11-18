package main

import "fmt"

func ubah(x *int) {
	*x = *x + 5
}

func main() {
	n := 10
	ubah(&n)
	fmt.Println(n)
}
