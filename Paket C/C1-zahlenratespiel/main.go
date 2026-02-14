package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Zufallszahl zwischen 0 und 100 erzeugen
	gesuchteZahl := rand.Intn(101)

	fmt.Println("Versuchen Sie eine natürliche Zahl zu erraten. Sie liegt zwischen 0 und 100.")

	scanner := bufio.NewScanner(os.Stdin)
	var versuch int
	var vorherigerAbstand int

	for versuchNummer := 1; ; versuchNummer++ {
		fmt.Printf("Ihr %d. Versuch:\n", versuchNummer)

		// Eingabe lesen
		if !scanner.Scan() {
			break
		}

		eingabe := strings.TrimSpace(scanner.Text())
		var err error
		versuch, err = strconv.Atoi(eingabe)

		if err != nil {
			fmt.Println("Bitte geben Sie eine ganze Zahl ein.")
			versuchNummer--
			continue
		}

		// Abstand berechnen
		abstand := versuch - gesuchteZahl
		if abstand < 0 {
			abstand = -abstand
		}

		// Rückmeldung geben
		if versuch == gesuchteZahl {
			fmt.Printf("Herzlichen Glückwunsch! Sie haben die Zahl in %d Versuchen erraten.\n", versuchNummer)
			break
		}

		// Switch-Anweisung für Rückmeldungen
		if versuchNummer == 1 {
			switch {
			case abstand <= 5:
				fmt.Println("heiß")
			default:
				fmt.Println("kalt")
			}
			vorherigerAbstand = abstand
		} else {
			switch {
			case abstand <= 5:
				fmt.Println("heiß")
			case abstand < vorherigerAbstand:
				fmt.Println("wärmer")
			case abstand > vorherigerAbstand:
				fmt.Println("kälter")
			default:
				fmt.Println("gleicher Abstand")
			}
			vorherigerAbstand = abstand
		}
	}
}
