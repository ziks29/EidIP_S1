package main

// bemerkenswert gibt die nächste natürliche Zahl größer start zurück, deren
// Quadrat die Summe der natürlichen Zahlen von 1 bis n ist (mit beliebigem n).
func bemerkenswert(start int) int {

	/* Hinweis zur Umsetzung:
	Da die Bedingung der for-Schleife leere ist, wird implizit true als Schleifenbedingung angenommen,
	d.h. es handelt sich um eine Endlosschleife.
	Laut Aufgabenstellung können wir allerdings davon ausgehen, dass die gesuchte Zahl existiert und
	somit die return-Anweisung innerhalb der Schleife zum Tragen kommen wird, wodurch ein Verlassen
	der Schleife quasi sichergestellt ist.
	*/

	// für alle Nachfolgezahlen
	for z := start + 1; ; z++ {
		// das Quadrat betrachten
		quadrat := z * z
		// Summe der Zahlen ab 1 berechnen
		summe := 0
		for i := 1; summe < quadrat; i++ {
			summe = summe + i
		}
		// wenn summe und quadrat gleich sind, ist die gesuchte Zahl gefunden.
		if summe == quadrat {
			return z
		}
		// ansonsten ist summe > quadrat und es muss die nächste Zahl geprüft werden ...
	}
}
