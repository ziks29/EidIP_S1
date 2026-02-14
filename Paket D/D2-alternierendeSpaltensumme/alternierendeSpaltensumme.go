package main

// alternierendeSpaltensumme überprüft arr daraufhin, ob in der letzten
// Zeile die alternierenden Spaltensummen stehen, korrigiert diese ggf.
// und gibt die Anzahl der fehlerhaften Summen zurück.
func alternierendeSpaltensumme(arr *[anzZeilen][anzSpalten]int) int {
	fehlerCount := 0
	lastRow := anzZeilen - 1

	// Für jede Spalte berechnen
	for spalte := 0; spalte < anzSpalten; spalte++ {
		// Alternierende Summe berechnen
		summe := 0
		for zeile := 0; zeile < lastRow; zeile++ {
			if zeile%2 == 0 {
				summe += arr[zeile][spalte]
			} else {
				summe -= arr[zeile][spalte]
			}
		}

		// Prüfen ob die letzte Zeile korrekt ist
		if arr[lastRow][spalte] != summe {
			fehlerCount++
			arr[lastRow][spalte] = summe
		}
	}

	return fehlerCount
}
