package main

import "math"

// rundreise gibt eine auf Basis der Entfernungsmatrix entfernung erstellte Route einer
// Rundreise und deren Gesamtstreckenlänge zurück.
func rundreise(entfernung [anzOrte][anzOrte]int) ([anzOrte + 1]int, int) {
	var route [anzOrte + 1]int
	var besucht [anzOrte]bool // vorbelegt mit false, noch kein Ort wurde besucht
	var nächste int
	gesamtstrecke := 0
	hier := 0
	besucht[hier] = true
	// Array-Indizes beginnen mit 0, aber die Orte sollen ab 1 nummeriert werden.
	// Somit wird "hier + 1" als erster Ort der Route eingetragen.
	route[0] = hier + 1
	orteIndex := 0

	// jetzt noch die restlichen anzOrte-1 Orte besuchen
	for besuchteOrte := 1; besuchteOrte < anzOrte; besuchteOrte++ {
		einzelstrecke := math.MaxInt
		// suche den nächstgelegenen, noch nicht besuchten Ort
		for spalte := 0; spalte < anzOrte; spalte++ {
			if !besucht[spalte] {
				if entfernung[hier][spalte] < einzelstrecke {
					einzelstrecke = entfernung[hier][spalte]
					nächste = spalte
				}
			}
		}
		gesamtstrecke = gesamtstrecke + einzelstrecke
		hier = nächste
		besucht[hier] = true
		route[orteIndex+1] = hier + 1
		orteIndex = orteIndex + 1
	}

	// und wieder zurück zum ersten Ort
	route[orteIndex+1] = 1
	gesamtstrecke = gesamtstrecke + entfernung[hier][0]

	return route, gesamtstrecke
}
