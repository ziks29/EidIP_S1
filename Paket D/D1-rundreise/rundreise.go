package main

// rundreise gibt eine auf Basis der Entfernungsmatrix entfernung erstellte Route einer
// Rundreise und deren Gesamtstreckenlänge zurück.
func rundreise(entfernung [anzOrte][anzOrte]int) ([anzOrte + 1]int, int) {
	var route [anzOrte + 1]int
	var besucht [anzOrte]bool
	
	// Starte bei Ort 1 (Index 0)
	aktuellerOrt := 0
	route[0] = 1 // Erste Position ist Ort 1
	besucht[0] = true
	gesamtstrecke := 0
	
	// Besuche alle weiteren Orte (zähler von 1 bis anzOrte-1)
	for zähler := 1; zähler < anzOrte; zähler++ {
		// Finde den nächstgelegenen unbesuchten Ort
		minEntfernung := int(^uint(0) >> 1) // Maximum int
		nächsterOrt := -1
		
		for ortIdx := 0; ortIdx < anzOrte; ortIdx++ {
			if !besucht[ortIdx] && entfernung[aktuellerOrt][ortIdx] < minEntfernung {
				minEntfernung = entfernung[aktuellerOrt][ortIdx]
				nächsterOrt = ortIdx
			}
		}
		
		// Gehe zum nächsten Ort
		gesamtstrecke += minEntfernung
		aktuellerOrt = nächsterOrt
		besucht[aktuellerOrt] = true
		route[zähler] = aktuellerOrt + 1 // Ortsnummern beginnen bei 1
	}
	
	// Kehre zu Ort 1 zurück
	gesamtstrecke += entfernung[aktuellerOrt][0]
	route[anzOrte] = 1
	
	return route, gesamtstrecke
}
