package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	fmt.Println(arr[0])

	// Membuat slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	// Menambahkan elemen ke slice
	slice = append(slice, 6, 7, 8)
	fmt.Println("Slice setelah ditambahkan elemen:", slice)

	// Mengakses elemen slice
	fmt.Println("slice[0]:", slice[0])
	fmt.Println("slice dari index 1 sampai index 4:", slice[1:4])
	fmt.Println("slice dari index 3 sampai akhir:", slice[3:])
	fmt.Println("slice dari index awal sampai index 2:", slice[:2])
	slice[1] = 10
	fmt.Println("slice[1] setelah diubah:", slice[1])

	// Menampilkan panjang dan kapasitas slice
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// Membuat slice dari array
	arr2 := [5]int{1, 2, 3, 4, 5}
	slice2 := arr2[1:4]
	fmt.Println("Slice dari array:", slice2)

	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source)) // Membuat slice tujuan dengan ukuran yang sama
	destination2 := make([]int, 3)          // Membuat slice tujuan dengan  3 elemen

	// Menyalin elemen dari source ke destination
	copy(destination, source)
	copy(destination2, source)

	fmt.Println("Source Slice: ", source)
	fmt.Println("Destination Slice: ", destination)
	fmt.Println("Destination2 Slice: ", destination2)

}
