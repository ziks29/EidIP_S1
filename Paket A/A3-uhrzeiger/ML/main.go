package main

import "fmt"

// Das Programm bestimmt den Winkel des Stunden- und des Minutenzeigers zu einer eingegebenen Uhrzeit.
func main() {
	var stunde, minute int
	fmt.Println("Geben Sie bitte eine Uhrzeit an, indem Sie zunächst ")
	fmt.Println("die Stunde (von 0 bis 23) und dann die Minute (von 0 bis 59) eingeben:")
	_, err := fmt.Scan(&stunde, &minute)

	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}

	// prüfen, ob die Eingaben eine gültige Uhrzeit darstellen
	if stunde < 0 || stunde > 23 || minute < 0 || minute > 59 {
		fmt.Printf("%v:%v ist keine gültige Uhrzeit.\n", stunde, minute)
		fmt.Println("Die Stunde muss im Bereich von 0 bis 23 und die Minute im Bereich von 0 bis 59 liegen.")
		return
	}

	// die Stellung der Zeiger wiederholt sich im 12-Stunden-Rhythmus
	stunde12 := stunde
	if stunde > 11 {
		stunde12 = stunde - 12
	}

	// Berechnung des Winkels des Minutenzeigers:
	// 60 Minuten verteilen sich auf 360°,
	// d.h. der Minutenzeiger bewegt sich um 6° pro Minute
	winkelMinutenZeiger := minute * 360 / 60

	// Berechnung des Winkels des Stundenzeigers:
	// 12 Stunden verteilen sich auf 360° und 60 Minuten auf den Winkelbereich einer Stunde,
	// d.h. der Stundenzeiger bewegt sich um 0.5 Grad pro Minute.
	// Also muss eine Typkonvertierung von int nach float64 vorgenommen werden,
	// damit auch die Nachkommastellen korrekt berechnet werden.
	winkelStundenZeiger := float64(stunde12)*360/12 + float64(minute)*360/12/60

	// Ausgabe der Ergebnisse
	fmt.Println()
	fmt.Printf("Zeigerstellung um %02d:%02d Uhr:\n", stunde, minute)
	fmt.Printf("Winkel des Stundenzeigers: %3.1f°\n", winkelStundenZeiger)
	fmt.Printf("Winkel des Minutenzeigers: %3d°\n", winkelMinutenZeiger)
}
