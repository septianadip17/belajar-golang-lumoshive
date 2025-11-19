package main

import "fmt"

func calcFee(hours int, member bool, holiday bool) int {
	if hours <= 0 {
		return 0
	}

	base := 5000
	extra := 0
	if hours > 2 {
		extra = (hours - 2) * 2000
	}

	total := base + extra

	if member {
		if hours <= 5 {
			total = total - total*50/100
		} else {
			total = total - total*30/100
		}
	}

	if holiday {
		total += 3000
	}

	return total
}

func main() {
	// Skenario 1: 4 jam, bukan member, bukan hari libur
	fee1 := calcFee(4, false, false)
	fmt.Println("Skenario 1 - 4 jam, bukan member, bukan hari libur")
	fmt.Println("Tarif parkir total: Rp", fee1)

	// Skenario 2: 2 jam, member, hari libur
	fee2 := calcFee(2, true, true)
	fmt.Println("Skenario 2 - 2 jam, member, hari libur")
	fmt.Println("Tarif parkir total: Rp", fee2)
}
