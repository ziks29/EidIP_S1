package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// Das Programm implementiert das Spiel Mastermind.
// Die Farben werden als Großbuchstaben A-F modelliert.
const farben = "ABCDEF"
const anzStellen = 4

func main() {
	const maxVersuche = 12

	var code, versuch [anzStellen]string
	for i := 0; i < anzStellen; i++ {
		zufall := rand.Intn(6)
		code[i] = farben[zufall : zufall+1]
	}

	versuchNr := 1
	erraten := false
	var eingabe string
	for (versuchNr <= maxVersuche) && !erraten {
		fmt.Printf("Dein %v. Rateversuch:\n", versuchNr)
		_, err := fmt.Scan(&eingabe)
		if err != nil {
			fmt.Println("Fehler: ", err)
			return
		}
		if len(eingabe) != anzStellen {
			fmt.Println("Fehler: Ihre Eingabe", eingabe, "hat keine 4 Stellen")
			return
		}
		for i := 0; i < anzStellen; i++ {
			f := eingabe[i : i+1]
			if strings.Index(farben, f) == -1 {
				fmt.Println("Fehler: Ihre Eingabe", eingabe, "enthält das Zeichen", f, "das keiner der Farben", farben, "entspricht")
				return
			}
			versuch[i] = f
		}
		schwarze, weiße := bewertung(versuch, code)
		erraten = schwarze == anzStellen
		if !erraten {
			fmt.Println(schwarze, "Schwarze und", weiße, "Weiße")
			versuchNr = versuchNr + 1
		}
		fmt.Println()
	}
	if erraten {
		fmt.Println("Gewonnen!!! :-)")
	} else {
		fmt.Println("Leider verloren :-(")
		fmt.Print("Der Code war ")
		for i := 0; i < anzStellen; i++ {
			fmt.Printf("%v", code[i])
		}
		fmt.Println()
	}
}
