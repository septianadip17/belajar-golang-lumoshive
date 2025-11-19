package main

import "fmt"

type Team struct {
	Name  string
	Score int
}

// hanya satu function selain main
func (t *Team) UpdateScore(newScore int) {
	if newScore > t.Score {
		t.Score = newScore
	}
}

func main() {
	team := Team{Name: "Garuda FC", Score: 2}

	fmt.Printf("Sebelum update -> Name: %s, Score: %d\n", team.Name, team.Score)

	team.UpdateScore(1) // tidak berubah karena 1 <= 2
	fmt.Printf("Setelah coba update 1 -> Name: %s, Score: %d\n", team.Name, team.Score)

	team.UpdateScore(5) // berubah karena 5 > 2
	fmt.Printf("Setelah coba update 5 -> Name: %s, Score: %d\n", team.Name, team.Score)
}
