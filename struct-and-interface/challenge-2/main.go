package main

import "fmt"

type Buku struct {
	Judul       string
	Penulis     string
	TahunTerbit int
}

func filterByYear(buku []Buku, minYear int) []Buku {
	var hasil []Buku
	for _, b := range buku {
		if b.TahunTerbit >= minYear {
			hasil = append(hasil, b)
		}
	}
	return hasil
}

func main() {
	buku := []Buku{
		{"Belajar Go", "Andi", 2018},
		{"Pemrograman Dasar", "Siti", 2012},
		{"Algoritma Praktis", "Budi", 2020},
		{"Web dengan Go", "Rina", 2015},
		{"Sistem Terdistribusi", "Tono", 2010},
		{"Machine Learning 101", "Agus", 2021},
		{"Jaringan Komputer", "Joko", 2014},
		{"Basis Data Modern", "Lina", 2016},
		{"Cloud Computing", "Nita", 2019},
		{"Arsitektur Software", "Hadi", 2013},
		{"Kecerdasan Buatan", "Arya", 2022},
		{"Struktur Data", "Eka", 2011},
		{"DevOps Handbook", "Rama", 2017},
		{"Pemrograman Mobile", "Sari", 2018},
		{"Keamanan Siber", "Fajar", 2020},
	}

	minTahun := 2015
	terfilter := filterByYear(buku, minTahun)

	fmt.Printf("Buku dengan tahun terbit >= %d:\n", minTahun)
	for _, b := range terfilter {
		fmt.Printf("- %s | %s | %d\n", b.Judul, b.Penulis, b.TahunTerbit)
	}
}
