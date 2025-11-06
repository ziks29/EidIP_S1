package main

// istPerfekt bestimmt, ob zahl eine perfekte Zahl ist.
// zahl muss eine natürliche Zahl sein.
func istPerfekt(zahl int) bool {
	kandidat := 0
	// mögliche perfekte Zahl berechnen,
	// indem alle Teiler von zahl zwischen 1 und zahl/2 aufsummiert werden
	for i := 1; i <= zahl/2; i++ {
		if zahl%i == 0 {
			kandidat = kandidat + i
		}
	}
	// checken, ob diese Summe mit zahl übereinstimmt
	return kandidat == zahl
}
