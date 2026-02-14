package main

import (
	"fmt"
	"strings"
)

// lesbarer gibt ihr Argument mit Unterstrichen statt Leerzeichen aus.
func lesbarer(s string) string {
	return strings.Replace(s, " ", "_", -1)
}

// Das Programm liest Nachrichten ein, verschlüsselt sie mit dem Skytale-Verfahren
// und versucht den Geheimtext zu knacken.
func main() {
	var nachricht string
	var scytale int
	fmt.Println("Welche Nachricht wollen Sie verschlüsseln?")
	// Anmerkung:
	// Mit den im Lerntext eingeführten Sprachmitteln ist es nicht möglich,
	// Strings mit Leerzeichen einzulesen. Daher wird hier darauf verzichtet.
	// Das Paket bufio kann dies hingegen leisten, ist aber nicht Thema des Lerntextes.
	// Leerzeichen in der Nachricht erleichtern zudem das Knacken der Botschaft,
	// insofern ist es auch sinnvoll, auf sie zu verzichten.
	fmt.Println("Bitte geben Sie nur Nachrichten ohne Leerzeichen ein.")
	_, err := fmt.Scan(&nachricht)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	fmt.Println("Wie ist der Durchmesser des Stabes in Buchstaben?")
	_, err = fmt.Scan(&scytale)
	if err != nil {
		fmt.Println("Fehler: ", err)
		return
	}
	if scytale < 2 {
		fmt.Println("Sie haben als Durchmesser", scytale, "eingegeben, sinnvoll ist aber nur ein Durchmesser größer 1.")
		return
	}

	// Wenn die letzte Zeile auf dem Stab nicht vollständig beschrieben wird,
	// wird die Nachricht um entsprechend viele Leerzeichen ergänzt.
	rest := len(nachricht) % scytale
	if rest > 0 {
		for i := rest; i < scytale; i++ {
			nachricht = nachricht + " "
		}
	}

	geheim := encrypt(nachricht, scytale)
	fmt.Println("Ihre Eingabe wird um Leerzeichen ergänzt, falls die letzte Zeile auf dem Stab nicht")
	fmt.Println("vollständig beschrieben wird.")
	fmt.Println("Leerzeichen werden als Unterstriche ausgegeben, um sie besser lesbar zu machen.")
	fmt.Println()
	fmt.Printf("«%v» wird verschlüsselt in «%v»\n", lesbarer(nachricht), lesbarer(geheim))
	fmt.Println()
	dmax := len(nachricht)
	fmt.Println("Wir versuchen die Verschlüsselung mit den Durchmessern 2 bis", dmax, "zu knacken")
	for d := 2; d <= dmax; d++ {
		fmt.Println("Durchmesser:", d, lesbarer(decrypt(geheim, d)))
	}
}
