package main

// istArmstrong bestimmt, ob zahl eine Armstrong-Zahl ist.
// zahl muss eine natürliche Zahl sein.
func istArmstrong(zahl int) bool {
	var rest, letzteZiffer int
	// Anzahl der Ziffern bestimmen
	// Wie oft lässt sich zahl durch 10 teilen?
	stellen := 1
	rest = zahl
	for rest >= 10 {
		stellen++
		rest = rest / 10
	}
	// mögliche Armstrong-Zahl berechnen
	kandidat := 0
	rest = zahl
	for rest > 0 {
		letzteZiffer = rest % 10
		kandidat = kandidat + potenziere(letzteZiffer, stellen)
		rest = rest / 10
	}
	// und checken
	return kandidat == zahl
}
