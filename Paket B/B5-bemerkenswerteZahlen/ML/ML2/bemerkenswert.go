package main

import "math"

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

	var summe int
	// Alle Dreieckszahlen durchgehen:
	// Die erste Quadratzahl, die größer ist als start*start, ist das Quadrat der gesuchten Zahl.
	for i := 1; ; i++ {
		summe = summe + i
		if summe > start*start {
			var wurzel float64 = math.Sqrt(float64(summe))
			// Durch die Typkonvertierung von float64 nach int werden die Nachkommastellen abgeschnitten,
			// d.h. die Typkonvertierung liefert ggf. einen abgerundeten Wurzel-Wert.
			var wurzelAbg int = int(wurzel)
			// Nur wenn summe eine Quadratzahl ist, ist die Wurzel von summe eine ganze Zahl.
			// In diesem Fall ändert die Konvertierung nichts am Wert,
			// sodass wurzelAbg und wurzel den gleichen Wert haben.
			if wurzel == float64(wurzelAbg) {
				return wurzelAbg
			}
			// ansonsten ist die aktuelle summe keine Quadratzahl
			// und es muss die nächste summe geprüft werden ...
		}
	}
}
