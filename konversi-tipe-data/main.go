package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 42
	var b int64 = int64(a) // Konversi dari int ke int64
	var c int8 = int8(a)   // Konversi dari int ke int8 (berhati-hatilah dengan kemungkinan overflow)

	fmt.Println(a, b, c)

	var x float32 = 3.14
	var y float64 = float64(x) // Konversi dari float32 ke float64
	var z float32 = float32(y) // Konversi dari float64 ke float32

	fmt.Println(x, y, z)

	var d int = 42
	var e float64 = float64(a) // Konversi dari int ke float64

	var f float64 = 3.14
	var g int = int(c) // Konversi dari float64 ke int (pecahan desimal akan dibuang)

	fmt.Println(d, e, f, g)

	var str string = "Hello, Go!"
	var bytes []byte = []byte(str)  // Konversi dari string ke []byte
	var str2 string = string(bytes) // Konversi dari []byte ke string

	fmt.Println(str, bytes, str2)

	// conversion

	// Konversi dari string ke int
	var strs string = "123"
	var num, err = strconv.Atoi(strs)
	if err == nil {
		fmt.Println(num)
	}

	// Konversi dari int ke string
	var num3 int = 456
	var str3 string = strconv.Itoa(num3)
	fmt.Println(str3)

	// Konversi dari string ke float64
	var str5 string = "3.14"
	var H, err2 = strconv.ParseFloat(str5, 64)
	if err2 == nil {
		fmt.Println(H)
	}

	// Konversi dari float64 ke string
	var f2 float64 = 1.618
	var str4 string = strconv.FormatFloat(f2, 'f', 6, 64)
	fmt.Println(str4)

}
