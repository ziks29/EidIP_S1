package main

// alternierendeSpaltensumme überprüft arr daraufhin, ob in der letzten
// Zeile die alternierenden Spaltensummen stehen, korrigiert diese ggf.
// und gibt die Anzahl der fehlerhaften Summen zurück.
func alternierendeSpaltensumme(arr *[anzZeilen][anzSpalten]int) int {
	fehler := 0
	// für jede Spalte
	for spalte := 0; spalte < anzSpalten; spalte++ {
		sum := 0
		vorzeichen := 1
		// alternierende Spaltensumme bis zur vorletzten Zeile berechnen
		for zeile := 0; zeile < anzZeilen-1; zeile++ {
			sum = sum + vorzeichen*arr[zeile][spalte]
			vorzeichen = -vorzeichen
		}
		// Wert in der letzten Zeile ggf. korrigieren und den Fehler zählen
		if arr[anzZeilen-1][spalte] != sum {
			fehler++
			arr[anzZeilen-1][spalte] = sum
		}
	}
	return fehler
}
