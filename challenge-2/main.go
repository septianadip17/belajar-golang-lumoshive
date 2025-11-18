package main

import "fmt"

func main() {
	A := 30
	B := 24.5
	C := -45
	D := 0.67

	// ubah tipe int ke float64 agar bisa dihitung bersama
	E := (float64(A) + B) * float64(C) / D

	fmt.Println("nilai dari variabel E adalah", E)
}
