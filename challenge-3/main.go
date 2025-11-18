package main

import "fmt"

func main() {
	hargaJual := 150000.0
	hargaBeli := 100000.0
	biayaOperasional := 1000.0
	terjual := 100.0
	discount := 0.15

	hargaSetelahDiskon := hargaJual * (1 - discount)
	totalPendapatan := hargaSetelahDiskon * terjual
	totalBiaya := (hargaBeli * terjual) + (biayaOperasional * terjual)
	keuntungan := totalPendapatan - totalBiaya

	fmt.Printf("Harga Jual Setelah Diskon: Rp %.2f\n", hargaSetelahDiskon)
	fmt.Printf("Total Pendapatan: Rp %.2f\n", totalPendapatan)
	fmt.Printf("Total Biaya: Rp %.2f\n", totalBiaya)
	fmt.Printf("Total Keuntungan: Rp %.2f\n", keuntungan)
}
